package main

import (
	"flag"
	"fmt"
	"github.com/emildel/gopoll/frontend/templates"
	"net/http"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 81, "Frontend server port")

	flag.Parse()

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		templates.Healthcheck().Render(r.Context(), w)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}

}
