package database

import (
	"database/sql"
	"main/internal/app/movie/model"

	_ "github.com/mattn/go-sqlite3"
)

func StartConnection(DBName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	model.Migrate(db)

	return db, nil
}
