package usecase

import (
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/dataprovider"
	"github.com/raa07/bolt20/entity"
)

type MainPage struct {
	db database.Database
}

func NewMainPage(db database.Database) (MainPage, error) {
	return MainPage{db}, nil
}

func (mp MainPage) CoursesList() ([]entity.Course, error) {
	cdp, err := dataprovider.NewCourseDataProvider(mp.db)
	if err != nil {
		return []entity.Course{}, err
	}

	courseList, err := cdp.AllCourses()
	if err != nil {
		return []entity.Course{}, err
	}

	return courseList, nil
}

//сделать через фабрику/registry и то же самое с dataprovider
// snake case в названиях файлов
// однобуквенные для ресиверов, маленькие функции - обычно норм названия
// упакавать темплейты в бинарник
// usecase возращает в дто и контроллер парсит в свою дто
