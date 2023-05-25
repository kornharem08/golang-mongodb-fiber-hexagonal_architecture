package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Title string             `bson:"title" json:"title"`
}

type MovieRepository interface {
	GetAll(page, limit int) ([]Movie, error)
}
