package app

import (
	"context"
	"log"
	"os"

	"github.com/swaggest/usecase"

	"web-hello/internal/db"
)

type Config struct {
	Env  string
	Dsn  string
	Port int
}

type application struct {
	Logger  *log.Logger
	Version string
	Config  Config
}

var app application

func InitApp(cfg Config, version string) application {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app = application{
		Config:  cfg,
		Logger:  logger,
		Version: version,
		// models: data.NewModels(db.GetDB()),
	}

	return app
}

func GetApp() application {
	return app
}

// Controller
func healthcheck() usecase.Interactor {
	type checkState struct {
		Status      string `json:"status"`
		DbPing      string `json:"db_ping"`
		Environment string `json:"environment"`
		Version     string `json:"version"`
	}

	u := usecase.NewInteractor(func(_ context.Context, _ struct{}, output *checkState) error {
		var ping string

		conn := db.GetDB()
		err := conn.Ping()
		if err != nil {
			ping = "no-ok"
		} else {
			ping = "ok"
		}

		// app := app.GetApp()

		data := checkState{
			Status:      "available",
			DbPing:      ping,
			Environment: app.Config.Env,
			Version:     app.Version,
		}

		*output = data
		return nil
	})

	u.SetTags("Health Check")
	return u
}
