package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error(err.Error(), "request_method", r.Method, "request_url", r.URL.String())
}

// This method gets called to output JSON to the webpage, don't need this...yet
//func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
//	err :=
//}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
}
