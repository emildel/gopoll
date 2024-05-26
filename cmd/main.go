package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	ui "github.com/emildel/gopoll"
	"github.com/emildel/gopoll/internal/data"
	"github.com/go-playground/form/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/r3labs/sse/v2"
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
	models         data.Models
	sseServer      *sse.Server
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

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.NewWithCleanupInterval(dbpool, time.Hour*24)
	sessionManager.Cookie.Secure = true

	sseServer := sse.New()
	sseServer.AutoStream = true

	app := &application{
		config:         cfg,
		logger:         logger,
		formDecoder:    formDecoder,
		sessionManager: scs.New(),
		models:         data.NewModel(dbpool),
		sseServer:      sseServer,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    app.newTlSConfig(),
	}

	logger.Info("starting server", "addr", server.Addr, "env", cfg.env)

	// Passing in empty params to ListenAndServeTLS since our certificate files are included in tlsConfig
	if app.config.env != "production" {
		err = server.ListenAndServeTLS("", "")
	}

	// If we are running on production, we will use Caddy to manage our TLS certs, so
	// can simply run the server without passing in any certs
	err = server.ListenAndServe()
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

func (app *application) newTlSConfig() *tls.Config {
	dir, err := ui.TLS.ReadDir("tls")
	if err != nil {
		app.logger.Error("error reading directory", "error", err)
	}

	certSlice := make([][]byte, 2)

	for index, file := range dir {
		readFile, err := ui.TLS.ReadFile(fmt.Sprintf("tls/%s", file.Name()))
		if err != nil {
			app.logger.Error("error reading file", "error", err)
		}
		certSlice[index] = readFile
	}

	certs, err := tls.X509KeyPair(certSlice[0], certSlice[1])
	if err != nil {
		app.logger.Error("error loading key pair", "error", err)
	}

	return &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		Certificates:     []tls.Certificate{certs},
	}
}
