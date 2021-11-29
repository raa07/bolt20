package dataprovider

import (
	"strconv"

	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/entity"
)

type LessonDataProvider struct {
	db database.Database
}

func NewLessonDataProvider(db database.Database) (LessonDataProvider, error) {
	return LessonDataProvider{db}, nil
}

func (p LessonDataProvider) LessonsByModuleId(id uint64) ([]entity.Lesson, error) {
	lessons := []entity.Lesson{}
	err := p.db.Connection.Select(&lessons, "SELECT * FROM lesson WHERE id_module=?;", strconv.FormatUint(id, 10))

	if err != nil {
		return []entity.Lesson{}, err
	}

	return lessons, nil
}

func (p LessonDataProvider) LessonById(id uint64) (entity.Lesson, error) {
	lesson := entity.Lesson{}
	err := p.db.Connection.Get(&lesson, "SELECT * FROM lesson WHERE id=?;", strconv.FormatUint(id, 10))
	if err != nil {
		return lesson, err
	}

	return lesson, nil
}
