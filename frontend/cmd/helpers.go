package main

import (
	"errors"
	"github.com/go-playground/form/v4"
	"net/http"
)

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
