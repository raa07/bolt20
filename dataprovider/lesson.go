package dataprovider

import (
	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/entity"
	"strconv"
)

type LessonDataProvider struct {
	db database.Database
}

func NewLessonDataProvider(db database.Database) (LessonDataProvider, error) {
	return LessonDataProvider{db}, nil
}

func (p LessonDataProvider) GetLessonsByModuleId(id uint64) ([]entity.Lesson, error) {
	lessons := []entity.Lesson{}
	err := p.db.Connection.Select(&lessons, "SELECT * FROM lesson WHERE id_module=?;", strconv.FormatUint(id, 10))

	if err != nil {
		return []entity.Lesson{}, err
	}

	return lessons, nil
}
