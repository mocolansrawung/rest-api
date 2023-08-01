package service

import (
	"database/sql"
	"errors"
	"fmt"
	"main/internal/app/movie/model"

	"github.com/gofrs/uuid"
)

type MovieService struct {
	DB *sql.DB
}

func NewMovieService(db *sql.DB) *MovieService {
	return &MovieService{
		DB: db,
	}
}

// Create a Movie
func (s *MovieService) CreateMovie(movie *model.Movie) error {
	query := `INSERT INTO movies (id, title, description, genre, imageUrl) VALUES (?, ?, ?)`

	id, _ := uuid.NewV4()
	_, err := s.DB.Exec(query, id, movie.Title, movie.Description, movie.Genre, movie.ImageUrl)

	return err
}

// Retrieve All Movies
func (s *MovieService) GetMoviesByFilter(genre string, sort string, order string, limit int, page int) ([]model.Movie, error) {
	var movies []model.Movie

	query := "SELECT * FROM movies WHERE 1=1"

	var args []interface{}

	if genre != "" {
		query += " AND genre = ?"
		args = append(args, genre)
	}

	if sort != "" {
		validColumns := map[string]bool{
			"id":          true,
			"title":       true,
			"description": false,
			"genre":       true,
			"imageUrl":    false,
		}
		if !validColumns[sort] {
			return nil, errors.New("Invalid sort parameter")
		}

		validOrders := map[string]bool{
			"asc":  true,
			"desc": true,
		}
		if !validOrders[order] {
			return nil, errors.New("Invalid order parameter")
		}

		query += fmt.Sprintf(" ORDER BY %s %s", sort, order)
	}

	offset := page * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie model.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Genre, &movie.ImageUrl)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

// Update a Movie
func (s *MovieService) UpdateMovie(id string, movie *model.Movie) error {
	query := `UPDATE movies SET title=?, description=?, genre=?, imageUrl=? WHERE id=?`

	_, err := s.DB.Exec(query, movie.Title, movie.Description, movie.Genre, movie.ImageUrl, id)

	return err
}

// Delete a Movie
func (s *MovieService) DeleteMovie(id string) error {
	query := `DELETE FROM movies WHERE id = ?`
	_, err := s.DB.Exec(query, id)

	return err
}
