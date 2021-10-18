package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/usecase"
	"log"
	"net/http"
)

type PublicController struct {
	db   database.Database
	echo *echo.Echo
}

//add main page request
//add main page view data for template

func NewPublicController(db database.Database, echo *echo.Echo) (PublicController, error) {
	return PublicController{db, echo}, nil
}

func (pc PublicController) HandleMainPage(context echo.Context) error {
	mp, err := usecase.NewMainPage(pc.db)
	if err != nil {
		return err
	}
	cl, err := mp.GetCoursesList()
	log.Print(cl)

	return context.Render(http.StatusOK, "main_page.html", map[string]interface{}{})
}

/*


db reference - https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f
*/
