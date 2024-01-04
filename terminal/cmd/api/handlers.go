package main

import (
	"fmt"
	"io"
	"net/http"
)

// Make a GET request to Healthcheck endpoint of frontend, display the returned HTML contents.
func (app *application) testHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sessionID := query.Get("session")

	requestUrl := fmt.Sprintf("http://%s/test?session=%s", "localhost:81", sessionID)

	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	request.Header.Set("Content-Type", "text/html")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	b, err := io.ReadAll(response.Body)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.Write(b)

	defer response.Body.Close()
}

func (app *application) homepageHandler(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("http://%s/", "localhost:81")

	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	request.Header.Set("Content-Type", "text/html")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.Write(b)

	defer response.Body.Close()
}
