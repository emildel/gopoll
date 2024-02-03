package main

import (
	ui "github.com/emildel/gopoll/frontend"
	"github.com/emildel/gopoll/frontend/templates"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/assets/*filepath", fileServer)
	router.HandlerFunc(http.MethodGet, "/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		templates.Healthcheck().Render(r.Context(), w)
	})

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		templates.Homepage().Render(r.Context(), w)
	})

	router.HandlerFunc(http.MethodGet, "/joinSession", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		sessionID := query.Get("session")
		templates.JoinSession(sessionID).Render(r.Context(), w)
	})

	router.HandlerFunc(http.MethodGet, "/createSession", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("HX-Redirect", "/createSession")
		templates.CreateSession().Render(r.Context(), w)
	})

	router.HandlerFunc(http.MethodGet, "/addPollAnswer", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<input type="text" name="inputQuestion"/> <br />`))
	})

	router.HandlerFunc(http.MethodPost, "/createPoll", app.createPollPOST)

	return router

}
