package db

import (
	"context"
	"fmt"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func HaveRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("maxdev")
	collection := db.Collection("relations")

	condition := bson.M{
		"userId":             t.UserID,
		"userRelationshipId": t.UserRelationshipID,
	}
	var result models.Relationship
	err := collection.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println("Error db:", err)
		return false, err
	}

	return true, nil
}
