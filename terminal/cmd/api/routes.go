package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	//router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(fmt.Sprintf("Error page! Http 405, you tried to access this page with a %s request, which is invalid.", r.Method)))
	//})

	router.HandleMethodNotAllowed = false

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/", app.homepageHandler)

	router.HandlerFunc(http.MethodGet, "/v1/test", app.testHandler)

	return router
}
