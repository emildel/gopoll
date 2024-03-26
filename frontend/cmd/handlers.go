package main

import (
	"fmt"
	"github.com/emildel/gopoll/frontend/internal/data"
	"github.com/emildel/gopoll/frontend/templates"
	"github.com/patrickmn/go-cache"
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
	if sessionExists != "" {
		templates.JoinSession(poll.Title, poll.Answers, true).Render(r.Context(), w)
		return
	}

	templates.JoinSession(poll.Title, poll.Answers, false).Render(r.Context(), w)

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
		app.clientError(w, http.StatusInternalServerError)
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
		app.clientError(w, http.StatusInternalServerError)
		return
	}

	app.cacheManager.Set("PollId", pollId, cache.DefaultExpiration)
	app.cacheManager.Set("PollTitle", form.Title, cache.DefaultExpiration)

	templates.PollCreator(form.Title, form.Questions).Render(r.Context(), w)
}

func (app *application) answerPoll(w http.ResponseWriter, r *http.Request) {

	var answer data.AnswerPollForm
	pollId := strings.Split(r.Header.Get("Referer"), "=")[1]

	err := app.decodePostForm(r, &answer)
	if err != nil {
		app.clientError(w, http.StatusInternalServerError)
		return
	}

	// Postgres is 1 indexed (index starts at 1, not 0), so increase the index
	// value by 1 to update the result of the correct answer.
	err = app.models.Polls.Update(answer.Answer+1, pollId)
	if err != nil {
		app.clientError(w, http.StatusInternalServerError)
		return
	}

	templates.AnswerSubmitted().Render(r.Context(), w)
}
