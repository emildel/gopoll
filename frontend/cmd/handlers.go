package main

import (
	"fmt"
	"github.com/emildel/gopoll/frontend/templates"
	"github.com/patrickmn/go-cache"
	"net/http"
)

type createPollForm struct {
	Title     string   `form:"title"`
	Questions []string `form:"inputAnswer"`
}

func (app *application) joinSession(w http.ResponseWriter, r *http.Request) {
	session := r.URL.Query().Get("session")
	if session == "" {
		app.notFound(w)
		return
	}

	pollId, _ := app.cacheManager.Get("PollId")
	title, _ := app.cacheManager.Get("PollTitle")

	if session == pollId {
		templates.JoinSession(interfaceToString(session), interfaceToString(title)).Render(r.Context(), w)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (app *application) createPollPOST(w http.ResponseWriter, r *http.Request) {
	//err := r.ParseForm()
	//if err != nil {
	//	app.clientError(w, http.StatusBadRequest)
	//	return
	//}
	//
	//var form createPollForm
	//
	//// decode the values retrieved from form in POST request into createPollForm
	//err = app.decodePostForm(r, &form)
	//if err != nil {
	//	app.clientError(w, http.StatusBadRequest)
	//	return
	//}
	//

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

	app.cacheManager.Set("PollId", pollId, cache.DefaultExpiration)
	app.cacheManager.Set("PollTitle", form.Title, cache.DefaultExpiration)

	templates.PollCreator(form.Title, form.Questions).Render(r.Context(), w)
}
