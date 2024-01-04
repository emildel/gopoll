package main

import (
	"fmt"
	"io"
	"net/http"
)

// Make a GET requset to Healthcheck endpoint of frontend, display the returned HTML contents.
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("http://%s/healthcheck", "localhost:81")

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
