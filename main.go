package main

import (
	"log"
	"github.com/martinezga/categories-system-golang/api"
	"github.com/martinezga/categories-system-golang/config"
)

type App struct {
	Config config.Environment
}

func NewApp(config config.Environment) *App {
	return &App{
		Config: config,
	}
}

func (app *App) Run() error {
	// Initial configuration

	log.Print("Appp is running...")

	api.ServeApi()

	// Clean up or close connections

	log.Print("Closing app...")
	return nil
}

func main() {
	env := config.LoadEnvironment()
	// connect to the database

	// build HTTP server
	app := NewApp(env)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
