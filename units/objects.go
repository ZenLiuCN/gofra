package units

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JsonArray a json array base on []any.
type JsonArray []any

func (j JsonArray) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	if len(j) == 0 {
		return "[]", nil
	}
	return json.Marshal(j)
}
func (j JsonArray) Scan(src any) error {
	var raw []byte
	switch src := src.(type) {
	case string:
		raw = []byte(src)
	case []byte:
		raw = src
	default:
		return fmt.Errorf("type %T not supported by Scan", src)
	}
	return json.Unmarshal(raw, &j)
}

// JsonObject a json object base on map[string]any.
type JsonObject map[string]any

func (j JsonObject) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	if len(j) == 0 {
		return "{}", nil
	}
	return json.Marshal(j)
}

func (j JsonObject) Scan(src any) error {
	var raw []byte
	switch src := src.(type) {
	case string:
		raw = []byte(src)
	case []byte:
		raw = src
	default:
		return fmt.Errorf("type %T not supported by Scan", src)
	}
	return json.Unmarshal(raw, &j)
}

// Json a json container which T should any type but not an unaddressable or unserializable type.
// !Important this type must use as pointer.
type Json[T any] struct {
	V T
}

func (j *Json[T]) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j.V)
}

func (j *Json[T]) Scan(src any) error {
	var raw []byte
	switch src := src.(type) {
	case string:
		raw = []byte(src)
	case []byte:
		raw = src
	default:
		return fmt.Errorf("type %T not supported by Scan", src)
	}
	return json.Unmarshal(raw, &j.V)
}
