package main

import (
	"encoding/json"
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
	if sessionExists != "" {
		templates.JoinSession(poll.Title, poll.Answers, session, true).Render(r.Context(), w)
		return
	}

	templates.JoinSession(poll.Title, poll.Answers, session, false).Render(r.Context(), w)

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

	app.sessionChannel.createChannel(pollId)

	templates.PollCreator(form.Title, form.Questions, pollId).Render(r.Context(), w)
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

	app.sessionChannel.sendMessage(pollId, "Update")

	//go func() { app.notifierChannel <- "Update" }()

	templates.AnswerSubmitted().Render(r.Context(), w)
}

// Gets called by Client to update the chart. Function is blocked until app.notifierChannel
// receives an update. Once a new update is received, calls the database to get the updated results.
// app.notifierChannel is written to in the answerPoll() handler.
func (app *application) updateChart(w http.ResponseWriter, r *http.Request) {

	pollSessionFromUri := strings.Split(r.RequestURI, "/")[2]
	pollId := app.sessionManager.GetString(r.Context(), fmt.Sprintf("PollId%s", pollSessionFromUri))

	if pollId == "" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Block until channel receives an update
	app.sessionChannel.waitOnChannel(pollId)

	poll, err := app.models.Polls.Get(pollId)
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	jsonData := map[string]interface{}{
		"answers": poll.Answers,
		"results": poll.Results,
	}

	js, err := json.Marshal(jsonData)
	if err != nil {
		app.clientError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}
