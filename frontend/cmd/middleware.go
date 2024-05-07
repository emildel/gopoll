package main

import "net/http"

func (app *application) sseActivated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.sseServer.ServeHTTP(w, r)

		next.ServeHTTP(w, r)
	})
}
