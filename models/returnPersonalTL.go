package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReturnPersonalTL struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID             primitive.ObjectID `bson:"userId" json:"userId,omitempty"`
	UserRelationshipID primitive.ObjectID `bson:"userRelationshipId" json:"userRelationshipId,omitempty"`
	Tweet              struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
