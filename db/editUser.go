package db

import (
	"context"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func EditUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("maxdev")
	collection := db.Collection("users")

	data := make(map[string]interface{})

	if len(user.Name) > 0 {
		data["name"] = user.Name
	}
	if len(user.LastName) > 0 {
		data["lastName"] = user.LastName
	}
	data["birthDate"] = user.BirthDate
	if len(user.Avatar) > 0 {
		data["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		data["banner"] = user.Banner
	}
	if len(user.Bio) > 0 {
		data["bio"] = user.Bio
	}
	if len(user.Location) > 0 {
		data["location"] = user.Location
	}
	if len(user.WebSite) > 0 {
		data["webSite"] = user.WebSite
	}

	updateString := bson.M{
		"$set": data,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
