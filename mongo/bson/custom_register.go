package bson

import (
	"github.com/randiapr/yurago/mongo/decoder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

// GetRegister ..
func GetRegister() *bsoncodec.Registry {
	rb := bsoncodec.NewRegistryBuilder()
	var primitiveCodecs bson.PrimitiveCodecs
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	rb.RegisterDecoder(decoder.TTime, bsoncodec.ValueDecoderFunc(decoder.TimeDecodeValue))
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb.Build()
}
