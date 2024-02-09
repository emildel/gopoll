package main

import (
	"github.com/emildel/gopoll/frontend/templates"
	"net/http"
)

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

	//---------------------------------------------

	//bar := charts.NewBar()
	//
	//bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
	//	Title: "Results",
	//}))
	//
	//data := make([]opts.BarData, 0)
	//data = append(data, opts.BarData{Value: 5})
	//data = append(data, opts.BarData{Value: 2})
	//data = append(data, opts.BarData{Value: 4})
	//
	//bar.SetXAxis(form.Questions).
	//	AddSeries("Count", data, func(s *charts.SingleSeries) {
	//		charts.WithCircularStyleOpts(opts.CircularStyle{RotateLabel: false})
	//	})

	templates.Poll(form.Title, form.Questions).Render(r.Context(), w)
}
