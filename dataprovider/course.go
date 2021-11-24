package dataprovider

import (
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/entity"
	"strconv"
)

type CourseDataProvider struct {
	db database.Database
}

func NewCourseDataProvider(db database.Database) (CourseDataProvider, error) {
	return CourseDataProvider{db}, nil
}

func (p CourseDataProvider) GetAllCourses() ([]entity.Course, error) {
	var courses []entity.Course
	err := p.db.Connection.Select(&courses, "SELECT * FROM course")
	if err != nil {
		return courses, nil
	}

	return courses, nil
}

func (p CourseDataProvider) GetAllCoursesForUser(idUser uint) ([]entity.Course, error) {
	var courses []entity.Course
	err := p.db.Connection.Select(&courses, "TODO;", idUser)
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (p CourseDataProvider) GetCourseById(id uint64) (entity.Course, error) {
	course := entity.Course{}
	err := p.db.Connection.Get(&course, "SELECT * FROM course WHERE id=?;", strconv.FormatUint(id, 10))
	if err != nil {
		return course, err
	}

	return course, nil
}
