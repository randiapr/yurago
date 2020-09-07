package bson

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Unmarshal custom
func Unmarshal(b []byte, s interface{}) error {
	if err := bson.UnmarshalWithRegistry(GetRegister(), b, s); err != nil {
		log.Printf("ERROR: %s", err.Error())
		return err
	}
	return nil
}

// UnmarshalDecode from FindOne
func UnmarshalDecode(res *mongo.SingleResult, s interface{}) error {
	if res.Err() != nil {
		log.Printf("ERROR: %s", res.Err())
		return res.Err()
	}
	data, _ := res.DecodeBytes()
	if err := bson.UnmarshalWithRegistry(GetRegister(), data, s); err != nil {
		return err
	}
	return nil
}
