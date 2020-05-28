package db

import (
	"context"
	"time"
	"github.com/maximp14/golangreact/models"
)

func AddRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("maxdev")
	collection := db.Collection("relations")

	_, err := collection.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
