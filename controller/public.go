package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/usecase"
)

type Public struct {
	db   database.Database
	echo *echo.Echo
}

//add main page request
//add main page view data for template

func NewPublicController(db database.Database, echo *echo.Echo) (Public, error) {
	return Public{db, echo}, nil
}

func (pc Public) HandleMainPage(context echo.Context) error {
	mp, err := usecase.NewMainPage(pc.db)
	if err != nil {
		return err
	}
	cl, err := mp.GetCoursesList()
	log.Print(cl)

	return context.Render(http.StatusOK, "main_page", map[string]interface{}{
		"courses": cl,
	})
}

/*


db reference - https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f
*/
