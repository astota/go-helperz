package helpers

import (
	"fmt"
	"testing"

	"github.com/smartystreets/assertions/should"
)

type maxMinIn struct {
	A int
	B int
}

var minTests = []struct {
	in  maxMinIn
	out int
}{
	{maxMinIn{1, 2}, 1},
	{maxMinIn{2, 1}, 1},
}

var maxTests = []struct {
	in  maxMinIn
	out int
}{
	{maxMinIn{1, 2}, 2},
	{maxMinIn{2, 1}, 2},
}

type maxMin64In struct {
	A int64
	B int64
}

var min64Tests = []struct {
	in  maxMin64In
	out int64
}{
	{maxMin64In{1, 2}, 1},
	{maxMin64In{2, 1}, 1},
}

var max64Tests = []struct {
	in  maxMin64In
	out int64
}{
	{maxMin64In{1, 2}, 2},
	{maxMin64In{2, 1}, 2},
}

func TestMin(t *testing.T) {
	for _, tt := range minTests {
		t.Run(fmt.Sprintf("Check %d and %d", tt.in.A, tt.in.B), func(t *testing.T) {
			max := Min(tt.in.A, tt.in.B)
			if max != tt.out {
				t.Errorf("got %d, want %d", max, tt.out)
			}
		})
	}
}

func TestMax(t *testing.T) {
	for _, tt := range maxTests {
		t.Run(fmt.Sprintf("Check %d and %d", tt.in.A, tt.in.B), func(t *testing.T) {
			min := Max(tt.in.A, tt.in.B)
			if min != tt.out {
				t.Errorf("got %d, want %d", min, tt.out)
			}
		})
	}
}

func TestMin64(t *testing.T) {
	for _, tt := range min64Tests {
		t.Run(fmt.Sprintf("Check %d and %d", tt.in.A, tt.in.B), func(t *testing.T) {
			min := Min64(tt.in.A, tt.in.B)
			if min != tt.out {
				t.Errorf("got %d, want %d", min, tt.out)
			}
		})
	}
}

func TestMax64(t *testing.T) {
	for _, tt := range max64Tests {
		t.Run(fmt.Sprintf("Check %d and %d", tt.in.A, tt.in.B), func(t *testing.T) {
			max := Max64(tt.in.A, tt.in.B)
			if max != tt.out {
				t.Errorf("got %d, want %d", max, tt.out)
			}
		})
	}
}

type removeStringInSliceIn struct {
	Array  []string
	Values []string
}

var removeStringInSliceTests = []struct {
	in  removeStringInSliceIn
	out []string
}{
	{removeStringInSliceIn{[]string{"1", "2", "3", "4", "5"}, []string{"1", "3", "5"}}, []string{"2", "4"}},
	{removeStringInSliceIn{[]string{"1", "2", "3", "4", "5"}, []string{"5"}}, []string{"1", "2", "3", "4"}},
	{removeStringInSliceIn{[]string{"1", "2", "3", "4", "5"}, []string{""}}, []string{"1", "2", "3", "4", "5"}},
	{removeStringInSliceIn{[]string{"1", "2", "3", "4", "5"}, []string{"10", "4", "7", "3"}}, []string{"1", "2", "5"}},
	{removeStringInSliceIn{[]string{"1", "2", "2", "3", "3"}, []string{"2"}}, []string{"1", "3", "3"}},
	{removeStringInSliceIn{[]string{"2", "2", "2", "2", "2"}, []string{"2", "1"}}, []string{}},
	{removeStringInSliceIn{[]string{"2", "2", "2", "2", "2"}, []string{"5"}}, []string{"2", "2", "2", "2", "2"}},
	{removeStringInSliceIn{[]string{"1", "2", "3", "4", "5"}, []string{}}, []string{"1", "2", "3", "4", "5"}},
	{removeStringInSliceIn{[]string{}, []string{"1", "2", "3", "4", "5"}}, []string{}},
	{removeStringInSliceIn{[]string{"1", "2", "3", "4", "5"}, []string{"3", "3"}}, []string{"1", "2", "4", "5"}},
}

func TestRemoveStringFromSlice(t *testing.T) {
	for _, tt := range removeStringInSliceTests {
		t.Run(fmt.Sprintf("Try to remove %v from %v", tt.in.Values, tt.in.Array), func(t *testing.T) {
			res := RemoveStringFromSlice(tt.in.Array, tt.in.Values...)
			if should.Resemble(res, tt.out) != "" {
				t.Errorf("got %v, want %v", res, tt.out)
			}
		})
	}
}

func BenchmarkRemoveStringFromSlice20_3(b *testing.B) {
	array := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	values := []string{"5", "11", "16"}
	for n := 0; n < b.N; n++ {
		RemoveStringFromSlice(array, values...)
	}
}

func BenchmarkRemoveStringFromSlice40_11(b *testing.B) {
	array := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40"}
	values := []string{"5", "11", "16", "6", "7", "8", "9", "10", "21", "34", "37"}
	for n := 0; n < b.N; n++ {
		RemoveStringFromSlice(array, values...)
	}
}
func BenchmarkRemoveStringFromSlice40_20(b *testing.B) {
	array := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40"}
	values := []string{"5", "11", "16", "6", "7", "8", "9", "10", "21", "34", "37", "22", "23", "24", "25", "26", "27", "28", "29", "30"}
	for n := 0; n < b.N; n++ {
		RemoveStringFromSlice(array, values...)
	}
}

func BenchmarkRemoveStringFromSlice1000(b *testing.B) {
	b.StopTimer()
	n := 1000
	array := make([]string, 0, n)
	values := make([]string, 0, n)
	for i := 0; i < n; i++ {
		array = append(array, fmt.Sprintf("%d", i%100))
		values = append(values, fmt.Sprintf("%d", i%10))
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		RemoveStringFromSlice(array, values...)
	}
	b.StopTimer()
}

func BenchmarkRemoveStringFromSlice10000(b *testing.B) {
	b.StopTimer()
	n := 10000
	array := make([]string, 0, n)
	values := make([]string, 0, n)
	for i := 0; i < n; i++ {
		array = append(array, fmt.Sprintf("%d", i%100))
		values = append(values, fmt.Sprintf("%d", i%10))
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		RemoveStringFromSlice(array, values...)
	}
	b.StopTimer()
}

func BenchmarkRemoveStringFromSlice100000(b *testing.B) {
	b.StopTimer()
	n := 100000
	array := make([]string, 0, n)
	values := make([]string, 0, n)
	for i := 0; i < n; i++ {
		array = append(array, fmt.Sprintf("%d", i%100))
		values = append(values, fmt.Sprintf("%d", i%10))
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		RemoveStringFromSlice(array, values...)
	}
	b.StopTimer()
}
