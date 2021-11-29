package dataprovider

import (
	"strconv"

	"github.com/raa07/bolt20/database"
	"github.com/raa07/bolt20/entity"
)

type ModuleDataProvider struct {
	db database.Database
}

func NewModuleDataProvider(db database.Database) (ModuleDataProvider, error) {
	return ModuleDataProvider{db}, nil
}

func (p ModuleDataProvider) ModulesByCourseId(id uint64) ([]entity.Module, error) {
	courseId := strconv.FormatUint(id, 10)
	modules := []entity.Module{}
	err := p.db.Connection.Select(&modules, "SELECT * FROM module WHERE id_course=?;", courseId)

	if err != nil {
		return []entity.Module{}, err
	}

	return modules, nil
}

func (p ModuleDataProvider) ModuleById(id uint64) (entity.Module, error) {
	module := entity.Module{}
	err := p.db.Connection.Get(&module, "SELECT * FROM module WHERE id=?;", strconv.FormatUint(id, 10))
	if err != nil {
		return module, err
	}

	return module, nil
}
