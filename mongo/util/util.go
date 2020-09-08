package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Index set index
func Index(field string, isUnique bool) mongo.IndexModel {
	return mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(isUnique).SetName(field + "_idx"),
	}
}
