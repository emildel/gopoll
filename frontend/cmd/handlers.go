package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/emildel/gopoll/frontend/internal/data"
	"github.com/emildel/gopoll/frontend/templates"
	"github.com/r3labs/sse/v2"
	"net/http"
	"strings"
	"time"
)

func (app *application) joinSession(w http.ResponseWriter, r *http.Request) {
	session := r.URL.Query().Get("session")
	if session == "" {
		app.notFound(w)
		return
	}

	poll, err := app.models.Polls.Get(session)
	switch {
	case errors.Is(data.ErrRecordNotFound, err):
		//Poll with this sessionId does not exist, so return the form with error validation. No redirection
		html := fmt.Sprintf(`<div hx-target="this" hx-swap="outerHTML">
                        <h1 class="font-bold text-2xl mb-2">Join a session</h1>
                        <input type="text" id="joinSessionForm" class="p-5 border-2 border-solid border-red-400" name="session" value="%s" placeholder="Enter your session id..." autocomplete="off" /> <br />
                        <div class="text-red-400">Session does not exist</div>
						<input type="submit" value="Enter" class="mt-4 bg-slate-50 text-neutral-950 py-5 px-10 text-center duration-300 cursor-pointer border-2 border-solid border-slate-950 hover:bg-[#555555] hover:text-white" />
                    </div>`, session)

		w.Write([]byte(html))
		return
	}

	// No errors, so we will be redirecting to the joinPoll page
	w.Header().Set("HX-Redirect", fmt.Sprintf("/joinPoll?session=%s", session))

	// Try to get session and convert to string. Only care about the presence of a session, not its value
	_, pollCreator := app.sessionManager.Get(r.Context(), fmt.Sprintf("PollId%s", session)).(string)

	// If a cookie exists telling us the user created this poll, then render the
	// template showing poll results
	if pollCreator {
		templates.JoinSession(poll.Title, poll.Answers, poll.Results, session, true, app.config.env).Render(r.Context(), w)
		return
	}

	// Otherwise, the user did not create this poll, so render the template with the
	// possible answers for the user to answer
	templates.JoinSession(poll.Title, poll.Answers, poll.Results, session, false, app.config.env).Render(r.Context(), w)

}

func (app *application) createPollPOST(w http.ResponseWriter, r *http.Request) {
	sessionId := app.generateUniqueSessionId()

	app.sessionManager.Put(r.Context(), fmt.Sprintf("PollId%s", sessionId), sessionId)

	http.Redirect(w, r, fmt.Sprintf("/createPoll/%s", sessionId), http.StatusPermanentRedirect)
}

func (app *application) createPollPOSTWithSession(w http.ResponseWriter, r *http.Request) {

	pollSessionFromUri := strings.Split(r.RequestURI, "/")[2]
	pollId := app.sessionManager.GetString(r.Context(), fmt.Sprintf("PollId%s", pollSessionFromUri))

	if pollId == "" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var form data.CreatePollForm

	// decode the values retrieved from form in POST request into createPollForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	poll := &data.Poll{
		PollSession: pollId,
		Title:       form.Title,
		Answers:     form.Questions,
		Results:     make([]int, len(form.Questions)),
		ExpiresAt:   time.Now().Add(time.Minute * 60),
	}

	err = app.models.Polls.Insert(poll)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.sseServer.CreateStream(pollId)

	templates.PollCreator(form.Title, form.Questions, poll.Results, pollId, app.config.env).Render(r.Context(), w)
}

func (app *application) answerPoll(w http.ResponseWriter, r *http.Request) {

	var answer data.AnswerPollForm
	pollId := strings.Split(r.Header.Get("Referer"), "=")[1]

	err := app.decodePostForm(r, &answer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Postgres is 1 indexed (index starts at 1, not 0), so increase the index
	// value by 1 to update the result of the correct answer.
	poll, err := app.models.Polls.Update(answer.Answer+1, pollId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	jsonData := map[string]interface{}{
		"answers": poll.Answers,
		"results": poll.Results,
	}

	event, err := formatServerSentEvent(jsonData)
	if err != nil {
		fmt.Println(err)
	}

	app.sseServer.Publish(pollId, &sse.Event{
		Event: []byte(pollId),
		Data:  []byte(event),
	})

	w.WriteHeader(200)

	templates.AnswerSubmitted().Render(r.Context(), w)
}

func (app *application) createPollGETWithSession(w http.ResponseWriter, r *http.Request) {
	pollSessionFromUri := strings.Split(r.RequestURI, "/")[2]

	poll, err := app.models.Polls.Get(pollSessionFromUri)
	if err != nil {
		app.notFound(w)
		return
	}

	_, pollCreator := app.sessionManager.Get(r.Context(), fmt.Sprintf("PollId%s", pollSessionFromUri)).(string)

	if pollCreator {
		templates.JoinSession(poll.Title, poll.Answers, poll.Results, pollSessionFromUri, true, app.config.env).Render(r.Context(), w)
		return
	}

	// Otherwise, the user did not create this poll, so render the template with the
	// possible answers for the user to answer
	// No errors, so we will be redirecting to the joinPoll page
	//w.Header().Set("HX-Redirect", fmt.Sprintf("/joinPoll?session=%s", pollSessionFromUri))
	http.Redirect(w, r, fmt.Sprintf("/joinPoll?session=%s", pollSessionFromUri), http.StatusSeeOther)
	//templates.JoinSession(poll.Title, poll.Answers, poll.Results, pollSessionFromUri, false, app.config.env).Render(r.Context(), w)
}

func formatServerSentEvent(data any) (string, error) {
	m := map[string]any{
		"data": data,
	}

	buff := bytes.NewBuffer([]byte{})

	encoder := json.NewEncoder(buff)

	err := encoder.Encode(m)
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%v\n\n", buff.String()))

	return sb.String(), nil
}
