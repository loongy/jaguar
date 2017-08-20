package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

// Float64 is an float64 used for possibly null database columns.
type Float64 struct {
	sql.NullFloat64
}

// ValidFloat64 returns a new Float64 that is valid.
func ValidFloat64(value float64) Float64 {
	return Float64{sql.NullFloat64{Float64: value, Valid: true}}
}

// InvalidFloat64 returns a new Float64 that is valid.
func InvalidFloat64() Float64 {
	return Float64{sql.NullFloat64{Float64: 0, Valid: false}}
}

// MarshalJSON converts a Float64 to JSON. It returns the value if valid and null
// otherwise.
func (null Float64) MarshalJSON() ([]byte, error) {
	if null.Valid {
		return json.Marshal(null.Float64)
	}
	return []byte("null"), nil
}

// UnmarshalJSON converts JSON to a Float64. Nil input returns an invalid Float64,
// otherwise it returns a valid Float64.
func (null *Float64) UnmarshalJSON(data []byte) error {
	// Unmarshal
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch value.(type) {
	case int64:
		null.Float64 = float64(value.(int64))
		null.Valid = true
	case float64:
		null.Float64 = value.(float64)
		null.Valid = true
	case nil:
		null.Float64 = 0
		null.Valid = false
	default:
		return fmt.Errorf("Cannot unmarshal %v into NullFloat64", reflect.TypeOf(value).Name())
	}
	return nil
}
