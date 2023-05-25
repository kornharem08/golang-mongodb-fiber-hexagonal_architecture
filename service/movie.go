package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieResponse struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Title string             `bson:"title" json:"title"`
}

type MovieService interface {
	GetMovies(page, limit int) ([]MovieResponse, error)
}
