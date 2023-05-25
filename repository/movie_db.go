package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type movieRepositoryDB struct {
	collection *mongo.Collection
}

func NewMovieRepositoryDB(db *mongo.Database) MovieRepository {
	collection := db.Collection("movies")
	return &movieRepositoryDB{collection: collection}
}

func (r movieRepositoryDB) GetAll(page, limit int) ([]Movie, error) {
	var movies []Movie
	fmt.Println(page, limit)
	// Create a context
	ctx := context.TODO()

	skip := int64((page - 1) * limit)

	// Define options to customize the query
	options := options.Find()
	options.SetLimit(int64(limit))
	options.SetSkip(skip)
	// Execute the query
	cursor, err := r.collection.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}

	// Iterate through the cursor and decode each document into a Movie object
	for cursor.Next(ctx) {
		var movie Movie
		err := cursor.Decode(&movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)

	}
	fmt.Println(movies)
	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor
	cursor.Close(ctx)

	return movies, nil
}
