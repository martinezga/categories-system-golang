package main

import (
	"log"

	"github.com/martinezga/categories-system-golang/api"
	"github.com/martinezga/categories-system-golang/config"
	"github.com/martinezga/categories-system-golang/pkg/db"
	"gorm.io/gorm"
)

type App struct {
	Config config.Environment
	Db     *gorm.DB
}

func NewApp(config config.Environment, db *gorm.DB) *App {
	return &App{
		Config: config,
		Db:     db,
	}
}

func (app *App) Run() error {
	// Initial configuration

	log.Print("Appp is running...")

	api.ServeApi(app.Db)

	// Clean up or close connections

	log.Print("Closing app...")
	return nil
}

func main() {
	env := config.LoadEnvironment()
	// connect to the database
	DB := db.Init(env.DatabaseURL)

	// build HTTP server
	app := NewApp(env, DB)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
