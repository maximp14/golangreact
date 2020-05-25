package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"log"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("maxdev")
	collection := db.Collection("tweet")

	var result []*models.ReturnTweets

	condition := bson.M{
		"userId": ID,
	}

	opt := options.Find()
	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	opt.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, condition, opt)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var data models.ReturnTweets
		err := cursor.Decode(&data)
		if err != nil {
			return result, false
		}
		result = append(result, &data)
	}
	return result, true
}
