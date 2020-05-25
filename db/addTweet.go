package db

import (
	"context"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddTweet(t models.TweetPersist) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("maxdev")
	collection := db.Collection("tweet")

	data := bson.M{
		"userId":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := collection.InsertOne(ctx, data)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil

}
