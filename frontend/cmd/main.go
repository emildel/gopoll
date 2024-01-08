package main

import (
	"flag"
	"fmt"
	ui "github.com/emildel/gopoll/frontend"
	"github.com/emildel/gopoll/frontend/templates"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 81, "Frontend server port")

	flag.Parse()

	router := httprouter.New()

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/assets/*filepath", fileServer)

	router.HandlerFunc(http.MethodGet, "/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		templates.Healthcheck().Render(r.Context(), w)
	})

	fs := http.FileServer(http.Dir("../styles"))
	http.Handle("/", fs)

	//http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
	//	templates.Healthcheck().Render(r.Context(), w)
	//})

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Homepage().Render(r.Context(), w)
	})*/

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		session := query.Get("session")
		templates.Test(session).Render(r.Context(), w)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		panic(err)
	}

}
