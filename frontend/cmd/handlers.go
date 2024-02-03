package main

import "net/http"

type createPollForm struct {
	Title     string   `form:"title"`
	Questions []string `form:"inputAnswer"`
}

func (app *application) createPollPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form createPollForm

	// decode the values retrieved from form in POST request into createPollForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	w.Write([]byte(form.Title))

	for _, question := range form.Questions {
		w.Write([]byte(question))
	}
}
