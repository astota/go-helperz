package helpers

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

type TestStruct struct {
	IntField    int64
	TimeField   time.Time
	StringField string
	MapField    map[string]interface{}
	SliceField  []int64
}

func Test_StructMatcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Convey("Test wrong type", t, func() {
		matcher := StructMatcher{
			Matching: "not a struct, obviously",
		}
		passed := TestStruct{}
		So(matcher.Matches(passed), ShouldBeFalse)
	})
	Convey("Test StructMatcher to a sting", t, func() {
		matcher := StructMatcher{
			Matching: TestStruct{},
		}
		passed := "not a struct, obviously"
		So(matcher.Matches(passed), ShouldBeFalse)
	})

	Convey("Test StructMatcher", t, func() {
		matching := TestStruct{
			IntField:    1,
			TimeField:   time.Now(),
			StringField: "testString",
			MapField:    make(map[string]interface{}),
			SliceField:  make([]int64, 1),
		}

		Convey("Test StructMatcher success", func() {
			matcher := StructMatcher{
				Matching: matching,
			}
			passed := matching
			So(matcher.Matches(passed), ShouldBeTrue)
		})

		Convey("Test StructMatcher empty", func() {
			matcher := StructMatcher{}
			passed := TestStruct{
				IntField:    matching.IntField,
				TimeField:   time.Now(),
				StringField: matching.StringField,
				MapField:    matching.MapField,
				SliceField:  matching.SliceField,
			}
			So(matcher.Matches(passed), ShouldBeFalse)
		})
		Convey("Test StructMatcher wrong time", func() {
			matcher := StructMatcher{
				Matching: matching,
			}
			passed := TestStruct{
				IntField:    matching.IntField,
				TimeField:   time.Now(),
				StringField: matching.StringField,
				MapField:    matching.MapField,
				SliceField:  matching.SliceField,
			}
			So(matcher.Matches(passed), ShouldBeFalse)
		})
		Convey("Test StructMatcher with skip field", func() {
			matcher := StructMatcher{
				Matching:   matching,
				SkipFields: []string{"TimeField"},
			}
			passed := TestStruct{
				IntField:    matching.IntField,
				TimeField:   time.Now(),
				StringField: matching.StringField,
				MapField:    matching.MapField,
				SliceField:  matching.SliceField,
			}
			So(matcher.Matches(passed), ShouldBeTrue)
		})
		Convey("Test StructMatcher with matcher field", func() {
			matcher := StructMatcher{
				Matching:      matching,
				MatcherFields: map[string]gomock.Matcher{"TimeField": TimeMatcher{Matching: time.Now()}},
			}
			passed := TestStruct{
				IntField:    matching.IntField,
				TimeField:   time.Now(),
				StringField: matching.StringField,
				MapField:    matching.MapField,
				SliceField:  matching.SliceField,
			}
			So(matcher.Matches(passed), ShouldBeTrue)
		})
		Convey("Test StructMatcher with matcher field but with wrong matcher", func() {
			matcher := StructMatcher{
				Matching:      matching,
				MatcherFields: map[string]gomock.Matcher{"TimeField": TimeMatcher{Matching: time.Now().Add(time.Hour)}},
			}
			passed := TestStruct{
				IntField:    matching.IntField,
				TimeField:   time.Now(),
				StringField: matching.StringField,
				MapField:    matching.MapField,
				SliceField:  matching.SliceField,
			}
			So(matcher.Matches(passed), ShouldBeFalse)
		})
		Convey("Test StructMatcher to pointer", func() {
			matcher := StructMatcher{
				Matching: matching,
			}
			passed := &matching
			So(matcher.Matches(passed), ShouldBeFalse)
		})
		Convey("Test pointer in StructMatcher to pointer", func() {
			matcher := StructMatcher{
				Matching: &matching,
			}
			passed := &matching
			So(matcher.Matches(passed), ShouldBeTrue)
		})
		Convey("Test pointer in StructMatcher to pointer to pointer", func() {
			matcher := StructMatcher{
				Matching: &matching,
			}
			passed := &matching
			So(matcher.Matches(&passed), ShouldBeFalse)
		})
	})
}
