package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/jackc/pgx/v5/pgxpool"
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
	db struct {
		dsn             string
		maxOpenConns    int
		maxConnLifetime string
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
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GOPOLL_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.StringVar(&cfg.db.maxConnLifetime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dbpool, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
	}

	defer dbpool.Close()

	logger.Info("database connection pool established")

	formDecoder := form.NewDecoder()

	app := &application{
		config:         cfg,
		logger:         logger,
		formDecoder:    formDecoder,
		sessionManager: scs.New(),
		cacheManager:   cache.New(5*time.Minute, 10*time.Minute),
	}

	app.sessionManager.Store = pgxstore.New(dbpool)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", server.Addr, "env", cfg.env)

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), app.routes())
	if err != nil {
		panic(err)
	}
}

func openDB(cfg config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	dbpool.Config().MaxConns = int32(cfg.db.maxOpenConns)

	duration, err := time.ParseDuration(cfg.db.maxConnLifetime)
	if err != nil {
		return nil, err
	}
	dbpool.Config().MaxConnLifetime = duration

	return dbpool, nil
}
