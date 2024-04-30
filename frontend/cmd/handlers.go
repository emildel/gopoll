package main

import (
	"errors"
	"fmt"
	"github.com/emildel/gopoll/frontend/internal/data"
	"github.com/emildel/gopoll/frontend/templates"
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
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	sessionExists := app.sessionManager.GetString(r.Context(), fmt.Sprintf("PollId%s", session))
	// If a cookie exists telling us the user created this poll, then render the
	// template showing poll results
	if sessionExists != "" {

		// We're going to render the chart, so we need to get the latest results to make
		// sure the chart doesn't render empty

		templates.JoinSession(poll.Title, poll.Answers, poll.Results, session, true).Render(r.Context(), w)
		return
	}

	// Otherwise, the user did not create this poll, so render the template with the
	// possible answers for the user to answer
	templates.JoinSession(poll.Title, poll.Answers, poll.Results, session, false).Render(r.Context(), w)

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

	app.sessionChannel.CreateSubscription(pollId)

	templates.PollCreator(form.Title, form.Questions, poll.Results, pollId).Render(r.Context(), w)
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
	err = app.models.Polls.Update(answer.Answer+1, pollId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.sessionChannel.SendMessage(pollId, "Update")

	templates.AnswerSubmitted().Render(r.Context(), w)
}

// Gets called by Client to update the chart. Once a new update is received,
// calls the database to get the updated results. app.sessionChannel is written
// to in the answerPoll() handler.
func (app *application) updateChart(w http.ResponseWriter, r *http.Request) {

	pollSessionFromUri := strings.Split(r.RequestURI, "/")[2]
	pollId := app.sessionManager.GetString(r.Context(), fmt.Sprintf("PollId%s", pollSessionFromUri))

	if pollId == "" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Block until channel receives an update
	err := app.sessionChannel.WaitOnChannel(pollId)
	if err != nil {
		switch {
		case errors.Is(err, ErrBlockedChannel):
			w.Write([]byte("An error occurred processing your request"))
			return

		case errors.Is(err, ErrChannelNotFound):
			w.Write([]byte("Could not find a poll with that ID"))
			return
		default:
			w.Write([]byte(err.Error()))
			return
		}
	}

	poll, err := app.models.Polls.Get(pollId)
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	jsonData := map[string]interface{}{
		"answers": poll.Answers,
		"results": poll.Results,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pollResults": jsonData}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
