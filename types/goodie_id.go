package types

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"strconv"
	"strings"
)

type BMGID struct {
	uuid.UUID
}

var typBMGID = reflect.TypeOf(BMGID{})

var uuidVersion uuid.Version = 0xa

// JSON handling

// MarshalJSON will marshal BMGID backward compatible way into JSON format
func (id BMGID) MarshalJSON() ([]byte, error) {
	if id.Version() == uuidVersion {
		integer := binary.LittleEndian.Uint64(id.UUID[8:16])
		return []byte("\"" + strconv.FormatUint(integer, 10) + "\""), nil
	}
	return []byte("\"" + id.String() + "\""), nil
}

// UnmarshalJSON will unmarshal JSON instance of BMG ID or
// populate specific error if it cannot do so
func (id *BMGID) UnmarshalJSON(b []byte) error {
	cur := strings.Trim(string(b), `"`)
	if ui, uerr := uuid.Parse(cur); uerr == nil {
		id.UUID = ui
		return nil
	}

	integer, ierr := strconv.Atoi(cur)
	if ierr != nil {
		return &json.UnmarshalTypeError{
			Value:  cur,
			Type:   typBMGID,
			Struct: "BMGID",
			Field:  "uuid",
		}
	}

	bs := [16]byte{}
	bs[6] = byte(uuidVersion) << 4
	binary.LittleEndian.PutUint64(bs[8:16], uint64(integer))
	id.UUID = bs

	return nil
}

// Text handling

// MarshalText will marshal BMGID backward compatible way into text
func (id BMGID) MarshalText() ([]byte, error) {
	return id.UUID.MarshalText()
}

// UnmarshalText will unmarshal text presentation of BMG ID or
// populate specific error if it cannot do so
func (id *BMGID) UnmarshalText(b []byte) error {
	return id.UUID.UnmarshalText(b)
}

// SQL handling

// Value will create SQL insertable value from BMGID
func (id BMGID) Value() (driver.Value, error) {
	return id.UUID.Value()
}

// Scan unmarshal database value to BMGID
func (id *BMGID) Scan(value interface{}) error {
	return id.UUID.Scan(value)
}
