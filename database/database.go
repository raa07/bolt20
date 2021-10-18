package database

import (
	"github.com/jmoiron/sqlx"
)

type Config struct {
	DataSource string
}

type Database struct {
	Connection *sqlx.DB
}

func NewDatabase(config Config)(Database, error) {
	db, err := sqlx.Connect("mysql", config.DataSource)
	if err != nil {
		return Database{&sqlx.DB{}}, err
	}

	return Database{db}, nil
}
