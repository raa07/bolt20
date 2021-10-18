package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/raa07/bolt20/controllers"
	"github.com/raa07/bolt20/database"
	"log"
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

	db, err := database.NewDatabase(config.Database)

	logger.Info("Application started")
	logger.Info("Db connected: ", db)

	e := echo.New()

	err = InitRender(e)
	if err != nil {
		log.Fatal(err)
	}

	pc, err := controllers.NewPublicController(db, e)


	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/", pc.HandleMainPage)

	e.Logger.Fatal(e.Start(":80"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	panic(err)
}