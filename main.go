package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/raa07/bolt20/controller"
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/usecase"
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
/////////////////////////////
	useCase, err := usecase.NewCourse(db)

	pc, err := controller.NewPublicController(db, e)
	if err != nil {
		log.Fatal(err)
	}

	cc, err := controller.NewCourseController(e, useCase)
	if err != nil {
		log.Fatal(err)
	}

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/", pc.HandleMainPage)
	e.GET("/courses/:id", cc.CoursePage)
	e.GET("/courses/:idCourse/module/:idModule", cc.ModulePage)

	e.Logger.Fatal(e.Start(":80"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	//panic(err)
}