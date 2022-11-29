package executor

import (
	"encoding/json"
	"strings"
)

type TypeFlags int

const (
	TypeNull TypeFlags = iota
	TypeBool
	TypeInteger
	TypeFloat
	TypeString
)

func (t TypeFlags) String() string {
	switch t {
	case TypeNull:
		return "Null"
	case TypeBool:
		return "boolean"
	case TypeFloat:
		return "float"
	case TypeString:
		return "string"
	case TypeInteger:
		return "int"
	default:
		return "unknown type"
	}
}

func (t TypeFlags) IsNumber() bool {
	return t == TypeFloat || t == TypeInteger
}

func (t TypeFlags) IsString() bool {
	return t == TypeString
}

func (t TypeFlags) IsBool() bool {
	return t == TypeBool
}

func (t TypeFlags) IsNull() bool {
	return t == TypeNull
}

func getType(v interface{}) (interface{}, TypeFlags) {
	val := castFixedPoint(v)
	switch val.(type) {
	case int64:
		return val, TypeInteger
	case float64:
		return val, TypeFloat
	case string:
		return val, TypeString
	case bool:
		return val, TypeBool
	default:
		return v, TypeNull
	}
}

func castFixedPoint(value interface{}) interface{} {
	switch v := value.(type) {
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case uint:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case int:
		return int64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case json.Number:
		if strings.Contains(v.String(), ".") {
			fv, _ := v.Float64()
			return fv
		} else {
			iv, _ := v.Int64()
			return iv
		}
	}
	return value
}
