package types

import (
	"encoding/binary"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestTransformations(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 10000; i++ {
		testUUID(t)
		testInteger(t)
	}
}

func testUUID(t *testing.T) {
	t.Helper()

	ref, _ := uuid.NewRandom()
	id := BMGID{UUID: ref}

	bs, _ := id.MarshalJSON()
	tst := BMGID{}
	tst.UnmarshalJSON(bs)
	if tst.String() != ref.String() {
		t.Errorf("Expected: %s, got: %s", ref.String(), tst.String())
	}
}

func testInteger(t *testing.T) {
	t.Helper()

	ref := []byte("\"" + strconv.Itoa(int(rand.Int63())) + "\"")
	middle := BMGID{}
	middle.UnmarshalJSON(ref)
	tst, _ := middle.MarshalJSON()

	if string(tst) != string(ref) {
		t.Errorf("Expected: %s, got: %s", string(ref), string(tst))
	}
}

func BenchmarkUUIDMarshalJSON(b *testing.B) {
	ref, _ := uuid.NewRandom()
	id := BMGID{UUID: ref}

	for i := 0; i < b.N; i++ {
		id.MarshalJSON()
	}
}

func BenchmarkUUIDUnmarshalJSON(b *testing.B) {
	bs := []byte(`"6b568013-b949-486e-ae36-8146459d422a"`)
	id := BMGID{}
	for i := 0; i < b.N; i++ {
		id.UnmarshalJSON(bs)
	}
}

func BenchmarkIntegerMarshalJSON(b *testing.B) {
	id := BMGID{}
	bs := [16]byte{}
	bs[6] = 0xa0
	binary.LittleEndian.PutUint64(bs[8:16], uint64(1823671253762))
	id.UUID = bs

	for i := 0; i < b.N; i++ {
		id.MarshalJSON()
	}
}

func BenchmarkIntegerUnmarshalJSON(b *testing.B) {
	bs := []byte(`"12398329734564"`)
	id := BMGID{}
	for i := 0; i < b.N; i++ {
		id.UnmarshalJSON(bs)
	}
}

func TestMarshalText(t *testing.T) {
	cases := []struct {
		name string
		text string
	}{
		{"UUID4", "a1c87dd5-265a-499f-abf6-ad79008afa14"},
		{"integer id", "00000000-0000-a000-9e78-88cf4751ddb8"},
		{"nonstandard UUID version", "686874f7-54fd-967a-a7e2-5582e19d950f"},
	}

	for _, tst := range cases {
		t.Run(tst.name, func(t *testing.T) {
			u, _ := uuid.Parse(tst.text)
			id := BMGID{UUID: u}
			bs, _ := id.MarshalText()
			if tst.text != string(bs) {
				t.Errorf("expecterd '%s', got '%s'", tst.text, string(bs))
			}
		})
	}
}

func TestUnmarshalText(t *testing.T) {
	cases := []struct {
		name string
		text string
	}{
		{"UUID4", "a1c87dd5-265a-499f-abf6-ad79008afa14"},
		{"integer id", "00000000-0000-a000-9e78-88cf4751ddb8"},
		{"nonstandard UUID version", "686874f7-54fd-967a-a7e2-5582e19d950f"},
	}

	for _, tst := range cases {
		t.Run(tst.name, func(t *testing.T) {
			id := BMGID{}
			id.UnmarshalText([]byte(tst.text))
			if tst.text != id.String() {
				t.Errorf("expecterd '%s', got '%s'", tst.text, id.String())
			}
		})
	}
}

func BenchmarkMarshalText(b *testing.B) {
	u, _ := uuid.NewRandom()
	id := BMGID{UUID: u}
	for i := 0; i < b.N; i++ {
		id.MarshalText()
	}

}

func BenchmarkUnmarshalText(b *testing.B) {
	bs := []byte("fc5305e7-f1fd-47b9-9876-9a41432bdb70")
	id := BMGID{}
	for i := 0; i < b.N; i++ {
		id.UnmarshalText(bs)
	}
}

func TestValue(t *testing.T) {
	cases := []struct {
		name string
		text string
	}{
		{"UUID4", "a1c87dd5-265a-499f-abf6-ad79008afa14"},
		{"integer id", "00000000-0000-a000-9e78-88cf4751ddb8"},
		{"nonstandard UUID version", "686874f7-54fd-967a-a7e2-5582e19d950f"},
	}

	for _, tst := range cases {
		t.Run(tst.name, func(t *testing.T) {
			u, _ := uuid.Parse(tst.text)
			id := BMGID{UUID: u}
			bs, _ := id.Value()
			val, _ := bs.(string)
			if tst.text != val {
				t.Errorf("expecterd '%s', got '%s'", tst.text, val)
			}
		})
	}
}

func TestScan(t *testing.T) {
	cases := []struct {
		name string
		text string
	}{
		{"UUID4", "a1c87dd5-265a-499f-abf6-ad79008afa14"},
		{"integer id", "00000000-0000-a000-9e78-88cf4751ddb8"},
		{"nonstandard UUID version", "686874f7-54fd-967a-a7e2-5582e19d950f"},
	}

	for _, tst := range cases {
		t.Run(tst.name, func(t *testing.T) {
			id := BMGID{}
			id.Scan([]byte(tst.text))
			if tst.text != id.String() {
				t.Errorf("expecterd '%s', got '%s'", tst.text, id.String())
			}
		})
	}
}

func BenchmarkValue(b *testing.B) {
	u, _ := uuid.NewRandom()
	id := BMGID{UUID: u}

	for i := 0; i < b.N; i++ {
		id.Value()
	}
}

func BenchmarkScan(b *testing.B) {
	bs := []byte("3a3010b5-15e2-4ddd-bf27-7cd8ff5080eb")
	id := BMGID{}

	for i := 0; i < b.N; i++ {
		id.Scan(bs)
	}
}
