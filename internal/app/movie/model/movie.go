package model

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

type Movie struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	ImageUrl    string    `json:"imageUrl"`
}

func Migrate(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS movies (
			id PRIMARY KEY NOT NULL,
			title TEXT NOT NULL,
			description TEXT,
			genre TEXT,
			imageUrl TEXT
		);`,
	)

	if err != nil {
		panic(err)
	}
}
