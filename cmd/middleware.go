package main

import (
	"fmt"
	"github.com/emildel/gopoll/templates"
	"net/http"
	"strings"
)

func (app *application) alreadyAnsweredMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pollId := strings.Split(r.RequestURI, "=")[1]
		isAnswered := app.sessionManager.Exists(r.Context(), fmt.Sprintf("pollAnswered_%s", pollId))

		if isAnswered {
			w.Header().Set("HX-Redirect", "/answered")
			templates.PollAnswered().Render(r.Context(), w)
			return
		}

		next.ServeHTTP(w, r)
	}
}
