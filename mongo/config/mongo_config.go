package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	myDb *mongo.Database
)

// Init connection
func Init(uri, database string) {
	// init connection mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	myDb = client.Database(database)
}

// GetDatabase mongo
func GetDatabase() *mongo.Database {
	return myDb
}
