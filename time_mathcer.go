package helpers

import (
	"fmt"
	"time"
)

// TimeMatcher a struct for convenient time fields checking
type TimeMatcher struct {
	Matching interface{}
}

// String return a string value of Matching data
func (tm TimeMatcher) String() string {
	return fmt.Sprintf("match to time after %v", tm.Matching)
}

//Matches if input value is a time-based value and if it's more than in TimeMatcher then it will return true
func (tm TimeMatcher) Matches(x interface{}) bool {
	x, matching, success := toExample(x, tm.Matching, time.Time{})
	// the time should be more than Matcher's time
	if !success || x.(time.Time).Before(matching.(time.Time)) {
		return false
	}
	return true
}
