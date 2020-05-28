package db

import (
	"context"
	"time"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson"
)

func PersonalTL(ID string, page int) ([]models.ReturnPersonalTL, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("maxdev")
	collection := db.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userRelationshipId",
			"foreignField": "userId",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := collection.Aggregate(ctx, conditions)
	var result []models.ReturnPersonalTL
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
