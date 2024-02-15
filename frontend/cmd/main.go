package main

import (
	"flag"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/patrickmn/go-cache"
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
	config         config
	logger         *slog.Logger
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
	cacheManager   *cache.Cache
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

	formDecoder := form.NewDecoder()

	app := &application{
		config:         cfg,
		logger:         logger,
		formDecoder:    formDecoder,
		sessionManager: scs.New(),
		cacheManager:   cache.New(5*time.Minute, 10*time.Minute),
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

}
