package main

import (
	ui "github.com/emildel/gopoll/frontend"
	"github.com/emildel/gopoll/frontend/templates"
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

	router.Handler(http.MethodGet, "/joinPoll", dynamic.ThenFunc(app.joinSession))

	router.HandlerFunc(http.MethodGet, "/createPoll", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("HX-Redirect", "/createPoll")
		templates.CreateSession().Render(r.Context(), w)
	})

	router.Handler(http.MethodPost, "/createPoll", dynamic.ThenFunc(app.createPollPOST))
	router.Handler(http.MethodPost, "/createPoll/:sessionId", dynamic.ThenFunc(app.createPollPOSTWithSession))
	router.Handler(http.MethodGet, "/createPoll/:sessionId", dynamic.ThenFunc(app.createPollPOSTWithSession))

	return router

}
