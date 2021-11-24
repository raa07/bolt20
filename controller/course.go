package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/raa07/bolt20/usecase"
	"net/http"
	"strconv"
)

type CourseController struct {
	echo *echo.Echo
	useCase usecase.Course
}

func NewCourseController (e *echo.Echo, useCase usecase.Course) (CourseController, error) {
	return CourseController{e, useCase}, nil
}

func (c CourseController) CoursePage(context echo.Context) error {
	idCourseUint, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	coursePageData, err := c.useCase.CoursePage(idCourseUint)
	if err != nil {
		panic(err) //todo: fix panic
	}

	fmt.Println(coursePageData.Modules)

	return context.Render(http.StatusOK, "course_page", map[string]interface{}{
		"course": coursePageData.Course,
		"modules": coursePageData.Modules,
	})
}

func (c CourseController) ModulePage(context echo.Context) error {
	idModuleUint, err := strconv.ParseUint(context.Param("idModule"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	idCourseUint, err := strconv.ParseUint(context.Param("idCourse"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	modulePageData, err := c.useCase.ModulePage(idModuleUint, idCourseUint)
	if err != nil {
		panic(err) //todo: fix panic
	}

	fmt.Println(modulePageData)

	return context.Render(http.StatusOK, "module_page", map[string]interface{}{
		"lessons": modulePageData.Lessons,
		"module": modulePageData.Module,
		"course": modulePageData.Course,
	})
}
