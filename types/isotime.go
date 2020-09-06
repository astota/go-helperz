// Package types ISOTime definition
// Golang time formats do not include 2006-01-02T15:04:05Z
// To support the exact date-time ISO8601 format, this little type
// comes into work.
package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg/types"
)

// ISOTime is a replacement for Golang JSON RFC3339 Format
type ISOTime time.Time

// ISOTimeFormat One of the ISO-8601 Time Formats.
const ISOTimeFormat = time.RFC3339

// UnmarshalJSON for custom ISO Format
func (t *ISOTime) UnmarshalJSON(b []byte) (err error) {
	var jt time.Time
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*t = ISOTime(jt)
		return
	}
	jt, err = time.Parse(ISOTimeFormat, s)
	if err != nil {
		return err
	}
	*t = ISOTime(jt.UTC())
	return nil
}

// MarshalJSON for iso format
func (t ISOTime) MarshalJSON() ([]byte, error) {
	var nilTime = (time.Time{}).UnixNano()
	if time.Time(t).UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).UTC().Format(ISOTimeFormat))), nil
}

// Equal check for ISOTime
func (t ISOTime) Equal(target *ISOTime) bool {
	tx := time.Time(t)
	ta := time.Time(*target)
	return tx.Equal(ta)
}

// String type
func (t ISOTime) String() string {
	return time.Time(t).UTC().Format(ISOTimeFormat)
}

// Before is just a handy alias
func (t ISOTime) Before(t2 *ISOTime) bool {
	return time.Time(t).Before(time.Time(*t2))
}

// Value Definition for Golang SQL Driver Value
func (t ISOTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t)), nil
}

// Scan definition for Golang SQL Driver
func (t *ISOTime) Scan(b interface{}) error {
	if b == nil {
		*t = ISOTime(time.Time{})
		return nil
	}
	tm, err := types.ParseTime(b.([]byte))
	if err != nil {
		return err
	}
	*t = ISOTime(tm)
	return nil
}

// StrToISOTime - convert string to ISOTime return nil if not possible
func StrToISOTime(str string) (result *ISOTime, err error) {
	if str == "" {
		return
	}
	var t time.Time
	t, err = time.Parse(ISOTimeFormat, str)
	if err != nil {
		return
	}
	ti := ISOTime(t)
	result = &ti
	return
}

// Overlap checks the overlap condition between two date ranges defined by
// [x1 - x2] to [y1 - y2] and generate required messages to let the consumer
// know what is wrong with the dates.
func Overlap(x1, x2, y1, y2 *ISOTime) (check bool, fields, messages []string) {
	tx1 := time.Time(*x1)
	tx2 := time.Time(*x2)
	ty1 := time.Time(*y1)
	ty2 := time.Time(*y2)
	// Below control simply checks overlap from a mathematical approach instead of
	// a case control. So it does not give a brief information about the overlap
	check = (tx1.Before(ty2) && ty1.Before(tx2)) || (tx1.Equal(ty1)) || (tx2.Equal(ty2))

	fields = make([]string, 0)
	messages = make([]string, 0)
	// Find the overlap case for detailed error message
	// We are reporting for X. (X is the current index in a payload)

	if tx1.Equal(ty1) {
		fields = append(fields, "starts_at")
		messages = append(messages, "starts_at is equal to starts_at of target")
	}

	if tx2.Equal(ty2) {
		fields = append(fields, "ends_at")
		messages = append(messages, "ends_at is equal to ends_at of target")
	}

	if tx1.Before(ty1) {
		if tx2.Before(ty2) {
			fields = append(fields, "ends_at")
			messages = append(messages, "ends_at is between starts_at and ends_at of target")
		} else {
			fields = append(fields, "starts_at", "ends_at")
			messages = append(messages, "starts_at is before targets starts_at", "ends_at is after targets ends_at")
		}
	}
	if ty1.Before(tx1) {
		if tx2.Before(ty2) {
			fields = append(fields, "starts_at", "ends_at")
			messages = append(messages, "starts_at is after targets starts_at", "ends_at is before targets ends_at")
		} else {
			fields = append(fields, "starts_at")
			messages = append(messages, "starts_at is between starts_at and ends_at of target")
		}
	}

	return
}
