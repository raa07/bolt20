package usecase

import (
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/dataprovider"
	"github.com/raa07/bolt20/entity"
)

type Course struct {
	db                 database.Database
	CourseDataProvider dataprovider.CourseDataProvider
	ModuleDataProvider dataprovider.ModuleDataProvider
	LessonDataProvider dataprovider.LessonDataProvider
}

type CoursePage struct {
	Modules []entity.Module
	Course  entity.Course
}

type ModulePage struct {
	Lessons []entity.Lesson
	Module entity.Module
	Course  entity.Course
}

func NewCourse(db database.Database) (Course, error) {
	courseDataProvider, err := dataprovider.NewCourseDataProvider(db)
	if err != nil {
		return Course{}, err
	}
	moduleDataProvider, err := dataprovider.NewModuleDataProvider(db)
	if err != nil {
		return Course{}, err
	}
	lessonDataProvider, err := dataprovider.NewLessonDataProvider(db)
	if err != nil {
		return Course{}, err
	}

	return Course{
		db,
		courseDataProvider,
		moduleDataProvider,
		lessonDataProvider,
	}, nil
}

func (c Course) CoursePage(id uint64) (CoursePage, error) {
	course, err := c.CourseDataProvider.GetCourseById(id)
	if err != nil {
		return CoursePage{}, err
	}
	modules, err := c.ModuleDataProvider.GetModulesByCourseId(id)
	if err != nil {
		return CoursePage{}, err
	}

	return CoursePage{
		modules,
		course,
	}, nil
}

func (c Course) ModulePage(moduleId uint64, courseId uint64) (ModulePage, error) {
	lessons, err := c.LessonDataProvider.GetLessonsByModuleId(moduleId)
	if err != nil {
		return ModulePage{}, err
	}
	module, err := c.ModuleDataProvider.GetModuleById(moduleId)
	if err != nil {
		return ModulePage{}, err
	}

	course, err := c.CourseDataProvider.GetCourseById(courseId)
	if err != nil {
		return ModulePage{}, err
	}

	return ModulePage{
		lessons,
		module,
		course,
	}, nil
}
