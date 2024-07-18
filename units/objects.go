package units

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Json a json container which T should any type but not an unaddressable or unserializable type.
// !Important T type must not a pointer.
type Json[T any] struct {
	V     T    //the real value of T
	Valid bool // dose this value is nil
}

func (j *Json[T]) Set(v T) {
	j.V = v
	j.Valid = true
}
func (j Json[T]) Value() (driver.Value, error) {
	if !j.Valid {
		return emptyAny, nil
	}
	return json.Marshal(j.V)
}
func (j Json[T]) MarshalJSON() ([]byte, error) {
	if !j.Valid {
		return emptyAny, nil
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
	case nil:
		return nil
	default:
		return fmt.Errorf("type %T not supported by Scan", src)
	}
	j.Valid = true
	return json.Unmarshal(raw, &j.V)
}
func (j *Json[T]) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &j.V)
}

var emptyObject = []byte{'{', '}'}
var emptyArray = []byte{'[', ']'}
var emptyAny = []byte("null")

// JsonObject hold map as json object
type JsonObject struct {
	V map[string]any
}

func (j JsonObject) MarshalJSON() ([]byte, error) {
	if j.V == nil {
		return nil, nil
	}
	if len(j.V) == 0 {
		return emptyObject, nil
	}
	return json.Marshal(j.V)
}

func (j JsonObject) Value() (driver.Value, error) {
	if j.V == nil {
		return nil, nil
	}
	if len(j.V) == 0 {
		return "{}", nil
	}
	return json.Marshal(j.V)
}

func (j *JsonObject) Scan(src any) error {
	var raw []byte
	switch src := src.(type) {
	case string:
		raw = []byte(src)
	case []byte:
		raw = src
	case nil:
		return nil
	default:
		return fmt.Errorf("type %T not supported by Scan", src)
	}
	if j.V == nil {
		j.V = make(map[string]any)
	}
	return json.Unmarshal(raw, &j.V)
}
func (j *JsonObject) UnmarshalJSON(bytes []byte) error {
	if j.V == nil {
		j.V = make(map[string]any)
	}
	return json.Unmarshal(bytes, &j.V)
}

// JsonArray hold slice of any as json array
type JsonArray struct {
	V []any
}

func (j JsonArray) Value() (driver.Value, error) {
	if j.V == nil {
		return nil, nil
	}
	if len(j.V) == 0 {
		return "[]", nil
	}
	return json.Marshal(j.V)
}
func (j JsonArray) MarshalJSON() ([]byte, error) {
	if j.V == nil {
		return nil, nil
	}
	if len(j.V) == 0 {
		return emptyArray, nil
	}
	return json.Marshal(j.V)
}
func (j *JsonArray) Scan(src any) error {
	var raw []byte
	switch src := src.(type) {
	case string:
		raw = []byte(src)
	case []byte:
		raw = src
	case nil:
		return nil
	default:
		return fmt.Errorf("type %T not supported by Scan", src)
	}
	if j.V == nil {
		j.V = make([]any, 0)
	}
	return json.Unmarshal(raw, &j.V)
}
func (j *JsonArray) UnmarshalJSON(bytes []byte) error {
	if j.V == nil {
		j.V = make([]any, 0)
	}
	return json.Unmarshal(bytes, &j.V)
}

type ArrayJson[T any] struct {
	V []T
}

func (j ArrayJson[T]) Value() (driver.Value, error) {
	if j.V == nil {
		return nil, nil
	}
	if len(j.V) == 0 {
		return "[]", nil
	}
	return json.Marshal(j.V)
}
func (j ArrayJson[T]) MarshalJSON() ([]byte, error) {
	if j.V == nil {
		return nil, nil
	}
	if len(j.V) == 0 {
		return emptyArray, nil
	}
	return json.Marshal(j.V)
}
func (j *ArrayJson[T]) Scan(src any) error {
	var raw []byte
	switch src := src.(type) {
	case string:
		raw = []byte(src)
	case []byte:
		raw = src
	case nil:
		return nil
	default:
		return fmt.Errorf("type %T not supported by Scan", src)
	}
	if j.V == nil {
		j.V = make([]T, 0)
	}
	return json.Unmarshal(raw, &j.V)
}
func (j *ArrayJson[T]) UnmarshalJSON(bytes []byte) error {
	if j.V == nil {
		j.V = make([]T, 0)
	}
	return json.Unmarshal(bytes, &j.V)
}
