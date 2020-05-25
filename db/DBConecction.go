package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoConnect Connection Object*/
var MongoConnect = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://maximp14:ExtraDosE2@maxdev-ys3th.mongodb.net/test?retryWrites=true&w=majority")

/* ConnectDB connection fuction*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection OK")

	return client
}

/* CheckConnection ping to the db*/
func CheckConnection() bool {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
