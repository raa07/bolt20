package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/raa07/bolt20/controller"
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/usecase"
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
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Db connected: ", db)
	logger.Info("Application started")

	e := echo.New()

	err = InitRender(e)
	if err != nil {
		log.Fatal(err)
	}
	/////////////////////////////
	useCase, err := usecase.NewCourse(db)
	if err != nil {
		log.Fatal(err)
	}

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
	e.GET("/course/:id", cc.CoursePage)
	e.GET("/course/:idCourse/module/:idModule", cc.ModulePage)
	e.GET("/course/:idCourse/module/:idModule/lesson/:idLesson", cc.LessonPage)

	e.Logger.Fatal(e.Start(":80"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	//panic(err)
}
