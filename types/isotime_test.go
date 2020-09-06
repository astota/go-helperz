package types

import (
	"encoding/json"
	"testing"
	"time"
)

type testStructISOTime struct {
	Time ISOTime `json:"time"`
}

func TestMarshallingISOTime(t *testing.T) {
	atime, _ := time.Parse(ISOTimeFormat, "2018-01-01T00:00:00Z")

	test := testStructISOTime{
		Time: ISOTime(atime),
	}

	data, _ := json.Marshal(test)
	if string(data) != "{\"time\":\"2018-01-01T00:00:00Z\"}" {
		t.Fail()
	}

	var e testStructISOTime
	json.Unmarshal(data, &e)
	if e.Time != ISOTime(atime) {
		t.Fail()
	}
}

func TestOverlap(t *testing.T) {
	tx1, _ := StrToISOTime("2018-01-01T00:00:00Z")
	tx2, _ := StrToISOTime("2019-01-01T00:00:00Z")
	// case 1:
	// t1 ------------ t2
	//      t3 -- t4
	ty1, _ := StrToISOTime("2018-03-01T00:00:00Z")
	ty2, _ := StrToISOTime("2018-04-01T00:00:00Z")
	if check, _, _ := Overlap(tx1, tx2, ty1, ty2); !check {
		t.Log("case 1")
		t.Fail()
	}

	// case 2:
	// t1 ---------- t2
	//         t3 ------- t4
	ty2, _ = StrToISOTime("2019-04-01T00:00:00Z")
	ty1, _ = StrToISOTime("2018-08-01T00:00:00Z")
	if check, _, _ := Overlap(tx1, tx2, ty1, ty2); !check {
		t.Log("case 2")
		t.Fail()
	}

	// case 3
	//       t1 ----------- t2
	//  t3 -------- t4
	ty1, _ = StrToISOTime("2017-03-01T00:00:00Z")
	ty2, _ = StrToISOTime("2018-04-01T00:00:00Z")
	if check, _, _ := Overlap(tx1, tx2, ty1, ty2); !check {
		t.Log("case 3")
		t.Fail()
	}

	// case 4:
	//        t1 ---------- t2
	//  t3 ---------------------- t4
	ty1, _ = StrToISOTime("2017-03-01T00:00:00Z")
	ty2, _ = StrToISOTime("2019-04-01T00:00:00Z")
	if check, _, _ := Overlap(tx1, tx2, ty1, ty2); !check {
		t.Log("case 4")
		t.Fail()
	}

	// check non overlap
	// t1 ---------- t2
	//                  t3 ------- t4
	ty1, _ = StrToISOTime("2019-03-01T00:00:00Z")
	ty2, _ = StrToISOTime("2020-04-01T00:00:00Z")
	if check, _, _ := Overlap(tx1, tx2, ty1, ty2); check {
		t.Log("no case")
		t.Fail()
	}

}
