package main

import (
	"log"

	"github.com/raa07/bolt20/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		// Unrecoverable error.
		log.Fatal(err)
	}

	logger, err := NewLogger(config.Logger)
	if err != nil {
		// Unrecoverable error.
		log.Fatal(err)
	}

	logger.Info("Application started")

	e := echo.New()

	pc := controllers.PublicController{}

	e.GET("/", pc.HandleMainPage)

	e.Logger.Fatal(e.Start(":80"))
}
