package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"web-hello/internal/app"
	"web-hello/internal/db"
)

const version = "1.0.0"

func main() {
	var cmd app.Config

	flag.IntVar(&cmd.Port, "port", 4000, "API server port")
	flag.StringVar(&cmd.Env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cmd.Dsn, "db-dsn", os.Getenv("WEBHELLO_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	// logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := app.InitApp(cmd, version)

	conn, err := db.InitDB(cmd.Dsn)
	if err != nil {
		app.Logger.Fatal(err)
	}

	defer conn.Close()

	// starting service
	addr := fmt.Sprintf(":%d", cmd.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      app.Route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("starting %s server on %s", cmd.Env, addr)
	err = srv.ListenAndServe()
	app.Logger.Fatal(err)
}
