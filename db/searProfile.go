package db

import (
	"context"
	"fmt"
	"time"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoConnect.Database("maxdev")
	collection := db.Collection("users")
	defer cancel()

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Profile not found"+err.Error())
		return profile, err
	}
	return profile, nil
}