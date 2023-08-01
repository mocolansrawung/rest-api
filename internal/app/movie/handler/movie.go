package handler

import (
	"encoding/json"
	"main/internal/app/movie/model"
	"main/internal/app/movie/service"
	"main/tools"
	"net/http"

	"github.com/go-chi/chi"
)

type MovieHandler struct {
	Service *service.MovieService
}

func NewMovieHandler(s *service.MovieService) *MovieHandler {
	return &MovieHandler{
		Service: s,
	}
}

// CreateMovie godoc
// @Summary Create a new movie
// @Description Create a new movie with the input payload
// @Tags movies
// @Accept  json
// @Produce  json
// @Param movie body model.Movie true "Create movie"
// @Success 201 {object} model.Movie
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /v1/movies [post]
func (c *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		tools.JsonResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.Service.CreateMovie(&movie); err != nil {
		tools.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tools.JsonResponse(w, http.StatusCreated, "success", movie)
}

// @Summary Get movies
// @Description Get a list of movies
// @Tags movies
// @Produce  json
// @Param genre query string false "Genre of the movie"
// @Param sort query string false "Sort movies by field"
// @Param order query string false "Order of sorting"
// @Param limit query integer false "Limit the number of movies" default(10)
// @Param page query integer false "Page number for movies list" default(0)
// @Success 200 {array} model.Movie
// @Router /v1/movies [get]
func (c *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	genre := r.URL.Query().Get("genre")
	sort := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")

	limitStr := r.URL.Query().Get("limit")
	limit, err := tools.ConvertIdParamsToInt(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	pageStr := r.URL.Query().Get("page")
	page, err := tools.ConvertIdParamsToInt(pageStr)
	if err != nil || page < 0 {
		page = 0
	}

	movies, err := c.Service.GetMoviesByFilter(genre, sort, order, limit, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tools.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tools.JsonResponse(w, http.StatusOK, "success", movies)
}

// UpdateMovie godoc
// @Summary Update an existing movie
// @Description Update an existing movie with the input payload
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Param movie body model.Movie true "Update movie"
// @Success 200 {object} model.Movie
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /v1/movies/{id} [put]
func (c *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		tools.JsonResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id := chi.URLParam(r, "id")

	if err := c.Service.UpdateMovie(id, &movie); err != nil {
		tools.JsonResponse(w, http.StatusInternalServerError, err.Error(), movie)
		return
	}

	tools.JsonResponse(w, http.StatusOK, "success", movie)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie by its ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the movie to delete"
// @Success 200 {object} map[string]interface{} "Successfully deleted movie"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/movies/{id} [delete]
func (c *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := c.Service.DeleteMovie(id); err != nil {
		tools.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tools.JsonResponse(w, http.StatusOK, "success", nil)
}
