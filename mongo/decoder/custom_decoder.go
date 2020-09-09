package decoder

import (
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var (
	// TTime ..
	TTime = reflect.TypeOf(time.Time{})
)

// TimeDecodeValue to avoid UTC from default
func TimeDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != TTime {
		return bsoncodec.ValueDecoderError{Name: "TimeDecodeValue", Types: []reflect.Type{TTime}, Received: val}
	}

	var timeVal time.Time
	timeFormatString := "2006-01-02 15:04:05"
	switch vrType := vr.Type(); vrType {
	case bsontype.DateTime:
		dt, err := vr.ReadDateTime()
		if err != nil {
			return err
		}
		timeVal = time.Unix(dt/1000, dt%1000*1000000)
	case bsontype.String:
		// assume strings are in the isoTimeFormat
		timeStr, err := vr.ReadString()
		if err != nil {
			return err
		}
		timeVal, err = time.Parse(timeFormatString, timeStr)
		if err != nil {
			return err
		}
	case bsontype.Int64:
		i64, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		timeVal = time.Unix(i64/1000, i64%1000*1000000)
	case bsontype.Timestamp:
		t, _, err := vr.ReadTimestamp()
		if err != nil {
			return err
		}
		timeVal = time.Unix(int64(t), 0)
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a time.Time", vrType)
	}
	timeVal = timeVal.Local()
	val.Set(reflect.ValueOf(timeVal))
	return nil
}
