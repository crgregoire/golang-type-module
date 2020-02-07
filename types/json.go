package types

import (
	"bytes"
	"database/sql/driver"
	"errors"
)

//
// JSON used to support JSON in gorm
//
type JSON []byte

//
// Value returns the value of a JSON field in mysql
//
func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

//
// Scan reads a JSON byte slice
//
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

//
// MarshalJSON supports JSON Marshalling
//
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

//
// UnmarshalJSON supports JSON UnMarshalling
//
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

//
// IsNull checks if JSON is null
//
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

//
// Equals compares two JSON values
//
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
