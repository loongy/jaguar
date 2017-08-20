package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

// Bool is an bool used for possibly null database columns.
type Bool struct {
	sql.NullBool
}

// ValidBool returns a new Bool that is valid.
func ValidBool(value bool) Bool {
	return Bool{sql.NullBool{Bool: value, Valid: true}}
}

// InvalidBool returns a new Bool that is valid.
func InvalidBool() Bool {
	return Bool{sql.NullBool{Bool: false, Valid: false}}
}

// MarshalJSON converts a Bool to JSON. It returns the value if valid and null
// otherwise.
func (null Bool) MarshalJSON() ([]byte, error) {
	if null.Valid {
		return json.Marshal(null.Bool)
	}
	return []byte("null"), nil
}

// UnmarshalJSON converts JSON to a Bool. Nil input returns an invalid Bool,
// otherwise it returns a valid Bool.
func (null *Bool) UnmarshalJSON(data []byte) error {
	// Unmarshal
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch value.(type) {
	case bool:
		null.Bool = value.(bool)
		null.Valid = true
	case nil:
		null.Bool = false
		null.Valid = false
	default:
		return fmt.Errorf("Cannot unmarshal %v into NullBool", reflect.TypeOf(value).Name())
	}
	return nil
}
