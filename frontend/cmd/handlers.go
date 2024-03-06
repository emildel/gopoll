package main

import (
	"fmt"
	"github.com/emildel/gopoll/frontend/internal/data"
	"github.com/emildel/gopoll/frontend/templates"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

type createPollForm struct {
	Title     string   `form:"title"`
	Questions []string `form:"inputAnswer"`
}

func (app *application) joinSession(w http.ResponseWriter, r *http.Request) {
	session := r.URL.Query().Get("session")
	//if session == "" {
	//	app.notFound(w)
	//	return
	//}

	//pollId := app.sessionManager.GetString(r.Context(), "PollId")

	poll, err := app.models.Polls.Get(session)
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	if session == session {
		templates.JoinSession(interfaceToString(poll.PollSession), interfaceToString(poll.Title)).Render(r.Context(), w)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (app *application) createPollPOST(w http.ResponseWriter, r *http.Request) {
	sessionId := app.generateUniqueSessionId()

	app.sessionManager.Put(r.Context(), "PollId", sessionId)

	http.Redirect(w, r, fmt.Sprintf("/createPoll/%s", sessionId), http.StatusPermanentRedirect)
}

func (app *application) createPollPOSTWithSession(w http.ResponseWriter, r *http.Request) {

	pollId := app.sessionManager.GetString(r.Context(), "PollId")

	if pollId == "" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form createPollForm

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
