package main

import (
	"flag"
	"fmt"
	"net/http"
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
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API saerver port")
	flag.StringVar(&cfg.env, "environment", "test-dev", "Environment (test-dev | production)")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space seperated)", func(s string) error {
		cfg.cors.trustedOrigins = strings.Fields(s)
		return nil
	})
	flag.Parse()

	app := &application{
		config: cfg,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
