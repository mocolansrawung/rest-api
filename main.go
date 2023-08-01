package main

//go:generate go run github.com/swaggo/swag/cmd/swag init

import (
	"log"
	"main/api"
	"main/config"
	"main/internal/app/movie/handler"
	"main/internal/app/movie/service"
	"main/pkg/database"
)

func main() {
	config := config.NewConfig("config/config.json")

	db, err := database.StartConnection(config.Database.DBName)
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}
	defer db.Close()

	service := service.NewMovieService(db)
	handler := handler.NewMovieHandler(service)

	api.StartServer(config.Server.Port, handler)
}
