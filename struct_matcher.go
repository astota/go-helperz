package helpers

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

// StructMatcher can be used for structures checking, includes:
type StructMatcher struct {
	Matching      interface{}               // set of fields for checking
	SkipFields    []string                  // these fields will be ignored
	MatcherFields map[string]gomock.Matcher // should be checked via GoMock matcher
}

// String return a string value of matching fields
func (sm StructMatcher) String() string {
	return fmt.Sprintf("StructMatcher to %v", sm.Matching)
}

// Matches Match all fields of an input structure to all fields in the saved structure. if they are equal => return true
func (sm StructMatcher) Matches(verifiableData interface{}) bool {

	verifiableData, matching, success := toExampleKind(verifiableData, sm.Matching, reflect.Struct)
	if !success {
		return success
	}
	v := reflect.ValueOf(matching)
	passed := reflect.ValueOf(verifiableData)
	for i := 0; i < v.NumField(); i++ {
		if sm.SkipFields != nil {
			if StringInSlice(v.Type().Field(i).Name, sm.SkipFields) {
				continue
			}
		}
		if sm.MatcherFields != nil {
			if matcher, ok := sm.MatcherFields[v.Type().Field(i).Name]; ok {
				if matcher.Matches(passed.Field(i).Interface()) {
					continue
				} else {
					return false
				}
			}
		}

		if !reflect.DeepEqual(passed.Field(i).Interface(), v.Field(i).Interface()) {
			return false
		}

	}
	return true
}

// toType return convertibleData which has been converted to the "kind" type
func toType(convertibleData interface{}, kind interface{}) (changed interface{}, kinds []reflect.Kind, err error) {
	kinds = make([]reflect.Kind, 0)
	for {
		t := reflect.TypeOf(convertibleData)
		kinds = append(kinds, t.Kind())
		if t.Kind() == reflect.Ptr {
			convertibleData = reflect.ValueOf(convertibleData).Elem().Interface()
			continue
		}
		if t != reflect.TypeOf(kind) {
			err = errors.New("data type conversion error")
			return
		}
		break
	}

	changed = convertibleData
	return
}

// toExample try to cast a and b values to the same data type from an example
func toExample(a, b interface{}, example interface{}) (exampleA, exampleB interface{}, success bool) {

	if a == nil || b == nil {
		return
	}
	b, bKinds, err := toType(b, example)
	if err != nil {
		return
	}
	a, aKinds, err := toType(a, example)
	if err != nil {
		return
	}
	if len(bKinds) != len(aKinds) {
		return
	}
	success = true
	exampleA = a
	exampleB = b
	return
}

// toKind return convertibleData which has been converted to "kind"
func toKind(convertibleData interface{}, kind reflect.Kind) (changed interface{}, kinds []reflect.Kind, err error) {
	kinds = make([]reflect.Kind, 0)
	for {
		t := reflect.TypeOf(convertibleData)
		kinds = append(kinds, t.Kind())
		if t.Kind() == reflect.Ptr {
			convertibleData = reflect.ValueOf(convertibleData).Elem().Interface()
			continue
		}
		if t.Kind() != kind {
			fmt.Println()
			fmt.Println(t.Kind())
			err = errors.New("data type conversion error")
			return
		}
		break
	}

	changed = convertibleData
	return
}

// toExampleKind tyr to cast a and b values to the equal type reflect.Kind
func toExampleKind(a, b interface{}, example reflect.Kind) (exampleA, exampleB interface{}, success bool) {
	if a == nil || b == nil {
		return
	}
	b, bKinds, err := toKind(b, example)
	if err != nil {
		return
	}
	a, aKinds, err := toKind(a, example)
	if err != nil {
		return
	}
	if len(bKinds) != len(aKinds) {
		return
	}

	success = true
	exampleA = a
	exampleB = b
	return
}
