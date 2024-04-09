package main

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/form/v4"
	"math/rand"
	"net/http"
)

type envelope map[string]any

func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {

		// If we try to use an invalid target destination, the Decode() method will return
		// an error with the type *form.InvalidDecoderError. Use errors.As to check for this specific error
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}

		return err
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) generateUniqueSessionId() string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
