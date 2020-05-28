package db

import (
	"context"
	"fmt"
	"github.com/maximp14/golangreact/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func TimeLine(ID string, page int64, search string, category string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoConnect.Database("maxdev")
	collection := db.Collection("users")
	defer cancel()

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relationship
		r.UserID = ID
		r.UserRelationshipID = s.ID.Hex()

		include = false

		found, err = HaveRelationship(r)
		if category == "new" && found == false {
			include = true
		}
		if category == "follow" && found == true {
			include = true
		}
		if r.UserRelationshipID == ID {
			include = false
		}
		if include == true {
			s.Password = ""
			s.Bio = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)

	return results, true

}
