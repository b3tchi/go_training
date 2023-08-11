package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"web-hello/internal/data"
	"web-hello/internal/db"
)

const version = "1.0.0"

type config struct {
	env  string
	dsn  string
	port int
}

type application struct {
	logger *log.Logger
	models data.Models
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("WEBHELLO_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	conn, err := db.InitDB(cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	defer conn.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db.GetDB()),
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
