package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

// String is a string used for possibly null database columns.
type String struct {
	sql.NullString
}

// ValidString returns a new String that is valid.
func ValidString(value string) String {
	return String{sql.NullString{String: value, Valid: true}}
}

// InvalidString returns a new String that is invalid.
func InvalidString() String {
	return String{sql.NullString{String: "", Valid: false}}
}

// MarshalJSON converts a String to JSON. It returns the value if valid and
// null otherwise.
func (null String) MarshalJSON() ([]byte, error) {
	if null.Valid {
		return json.Marshal(null.String)
	}
	return []byte("null"), nil
}

// UnmarshalJSON converts JSON to a String. Nil input returns an invalid
// String, otherwise it returns a valid NullString.
func (null *String) UnmarshalJSON(data []byte) error {
	// Unmarshal
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch value.(type) {
	case string:
		null.String = value.(string)
		null.Valid = true
	case nil:
		null.String = ""
		null.Valid = false
	default:
		return fmt.Errorf("Cannot unmarshal %v into NullString", reflect.TypeOf(value).Name())
	}
	return nil
}
