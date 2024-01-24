package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

type config struct {
	port int
	env  string
	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 81, "Frontend server port")
	flag.StringVar(&cfg.env, "environment", "test-dev", "Environment (test-dev | production)")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space seperated)", func(s string) error {
		cfg.cors.trustedOrigins = strings.Fields(s)
		return nil
	})

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", server.Addr, "env", cfg.env)

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), app.routes())
	if err != nil {
		panic(err)
	}

	/*fileServer := http.FileServer(http.FS(ui.Files))
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

}
