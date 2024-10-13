package main

import (
	ui "github.com/emildel/gopoll"
	"github.com/emildel/gopoll/templates"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/assets/*filepath", fileServer)

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.HandlerFunc(http.MethodGet, "/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		templates.Healthcheck().Render(r.Context(), w)
	})

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		templates.Homepage().Render(r.Context(), w)
	})

	router.Handler(http.MethodGet, "/joinPoll", dynamic.Then(app.alreadyAnsweredMiddleware(app.joinSession)))

	// View the new poll creation page
	router.HandlerFunc(http.MethodGet, "/createPoll", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("HX-Redirect", "/createPoll")
		templates.CreateSession().Render(r.Context(), w)
	})

	// First POST called when 'Create poll' button is pressed, to generate a unique
	// ID to put in the URL
	router.Handler(http.MethodPost, "/createPoll", dynamic.ThenFunc(app.createPollPOST))

	// Redirect to new URL with the previously created unique ID
	router.Handler(http.MethodPost, "/createPoll/:sessionId", dynamic.ThenFunc(app.createPollPOSTWithSession))
	router.Handler(http.MethodGet, "/createPoll/:sessionId", dynamic.ThenFunc(app.createPollGETWithSession))

	router.Handler(http.MethodPost, "/answerPoll", dynamic.ThenFunc(app.answerPoll))

	//sseEnabled := dynamic.Append(app.sseActivated)

	router.HandlerFunc(http.MethodGet, "/updateChart/:sessionId", app.sseServer.ServeHTTP)

	router.HandlerFunc(http.MethodGet, "/answered", func(w http.ResponseWriter, r *http.Request) {
		templates.PollAnswered().Render(r.Context(), w)
	})

	return router
}
