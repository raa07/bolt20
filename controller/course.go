package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raa07/bolt20/usecase"
)

type Course struct {
	echo    *echo.Echo
	useCase usecase.Course
}

func NewCourseController(e *echo.Echo, useCase usecase.Course) (Course, error) {
	return Course{e, useCase}, nil
}

func (c Course) CoursePage(context echo.Context) error {
	idCourse, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	coursePageData, err := c.useCase.CoursePage(idCourse)
	if err != nil {
		panic(err) //todo: fix panic
	}

	fmt.Println(coursePageData.Modules)

	return context.Render(http.StatusOK, "course_page", map[string]interface{}{
		"course":  coursePageData.Course,
		"modules": coursePageData.Modules,
	})
}

func (c Course) ModulePage(context echo.Context) error {
	idModule, err := strconv.ParseUint(context.Param("idModule"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	idCourse, err := strconv.ParseUint(context.Param("idCourse"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	modulePageData, err := c.useCase.ModulePage(idModule, idCourse)
	if err != nil {
		panic(err) //todo: fix panic
	}

	return context.Render(http.StatusOK, "module_page", map[string]interface{}{
		"lessons": modulePageData.Lessons,
		"module":  modulePageData.Module,
		"course":  modulePageData.Course,
	})
}

func (c Course) LessonPage(context echo.Context) error {
	idModule, err := strconv.ParseUint(context.Param("idModule"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	idCourse, err := strconv.ParseUint(context.Param("idCourse"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	idLesson, err := strconv.ParseUint(context.Param("idLesson"), 10, 64)
	if err != nil {
		panic(err) //todo: fix panic
	}

	lessonPageData, err := c.useCase.LessonPage(idCourse, idModule, idLesson)
	if err != nil {
		panic(err) //todo: fix panic
	}

	return context.Render(http.StatusOK, "lesson_page", map[string]interface{}{
		"lesson": lessonPageData.Lesson,
		"module": lessonPageData.Module,
		"course": lessonPageData.Course,
	})
}
