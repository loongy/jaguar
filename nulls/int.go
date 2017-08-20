package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

// Int64 is an int64 used for possibly null database columns.
type Int64 struct {
	sql.NullInt64
}

// ValidInt64 returns a new Int64 that is valid.
func ValidInt64(value int64) Int64 {
	return Int64{sql.NullInt64{Int64: value, Valid: true}}
}

// InvalidInt64 returns a new Int64 that is valid.
func InvalidInt64() Int64 {
	return Int64{sql.NullInt64{Int64: 0, Valid: false}}
}

// MarshalJSON converts a Int64 to JSON. It returns the value if valid and null
// otherwise.
func (null Int64) MarshalJSON() ([]byte, error) {
	if null.Valid {
		return json.Marshal(null.Int64)
	}
	return []byte("null"), nil
}

// UnmarshalJSON converts JSON to a Int64. Nil input returns an invalid Int64,
// otherwise it returns a valid Int64.
func (null *Int64) UnmarshalJSON(data []byte) error {
	// Unmarshal
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch value.(type) {
	case int64:
		null.Int64 = value.(int64)
		null.Valid = true
	case float64:
		null.Int64 = int64(value.(float64))
		null.Valid = true
	case nil:
		null.Int64 = 0
		null.Valid = false
	default:
		return fmt.Errorf("Cannot unmarshal %v into NullInt64", reflect.TypeOf(value).Name())
	}
	return nil
}
