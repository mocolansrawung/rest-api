package api

import (
	"fmt"
	"main/docs"
	"main/internal/app/movie/handler"
	"main/internal/app/movie/middleware"
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func StartServer(port string, handler *handler.MovieHandler) {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.MockAuthMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	docs.SwaggerInfo.Title = "Movie Database"
	docs.SwaggerInfo.Version = "v1"
	conf := httpSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.Get("/swagger/*", httpSwagger.Handler(conf))

	r.Route("/v1/", func(r chi.Router) {
		r.Post("/movies", handler.CreateMovie)
		r.Get("/movies", handler.GetMovies)
		r.Route("/movies/{id}", func(r chi.Router) {
			r.Put("/", handler.UpdateMovie)
			r.Delete("/", handler.DeleteMovie)
		})
	})

	fmt.Printf("Listening and serving to Port %s\n", port)
	http.ListenAndServe(port, r)
}
