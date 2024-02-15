package main

import (
	"context"
	"net/http"
)

func (app *application) getSessionId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sessionId := app.sessionManager.GetString(r.Context(), "PollSessionId")

		if sessionId != "" {
			ctx := context.WithValue(r.Context(), sessionIdCtx, sessionId)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
