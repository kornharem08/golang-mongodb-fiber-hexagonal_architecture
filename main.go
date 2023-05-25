package main

import (
	"basic/database"
	"basic/handler"
	"basic/repository"
	"basic/service"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	client, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Close the MongoDB connection when you're done
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	app := fiber.New()
	db := client.Database("sample_mflix")
	movieRepositoryDB := repository.NewMovieRepositoryDB(db)
	movieService := service.NewMovieService(movieRepositoryDB)
	movieHandler := handler.NewMovieHandler(movieService)

	// Define routes
	app.Get("/movies", movieHandler.GetMovies)

	app.Listen(":8000")
}
