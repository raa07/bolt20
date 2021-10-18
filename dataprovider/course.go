package dataprovider

import (
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/entity"
)

type CourseDataProvider struct{
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

func (p CourseDataProvider) GetAllCoursesForUser() ([]entity.Course, error) {
	return []entity.Course{}, nil
}

