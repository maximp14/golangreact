package db

import (
	"context"
	"time"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("maxdev")
	collection := db.Collection("users")

	u.Password,_ = EncryptPassword(u.Password)

	result, err := collection.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil

}