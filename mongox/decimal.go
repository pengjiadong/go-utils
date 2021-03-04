package mongox

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MyFloat64Codec
// float64 <-> decimal128
type MyFloat64Codec struct {
}

// DecodeValue is the ValueDecoderFunc for time.Time.
func (*MyFloat64Codec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	var float64Val float64
	switch vrType := vr.Type(); vrType {
	case bsontype.Decimal128:
		v, err := vr.ReadDecimal128()
		if err != nil {
			return err
		}
		float64Val = cast.ToFloat64(v.String())
	case bsontype.Double:
		v, err := vr.ReadDouble()
		if err != nil {
			return err
		}
		float64Val = v
	}
	val.Set(reflect.ValueOf(float64Val))
	return nil
}

// EncodeValue is the ValueEncoderFunc for time.TIme.
func (codec *MyFloat64Codec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	float64Val := val.Interface().(float64)
	dd, _ := primitive.ParseDecimal128(Round(float64Val))
	return vw.WriteDecimal128(dd)
}

func Round(f float64) string {
	floatStr := fmt.Sprintf("%.2f", f)
	return floatStr
}
