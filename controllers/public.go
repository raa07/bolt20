package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PublicController struct {

}

func (pc PublicController) HandleMainPage(c echo.Context) error {
	return c.String(http.StatusOK, "Main page")
}