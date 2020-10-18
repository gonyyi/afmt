package afmt_test

import (
	"fmt"
	"github.com/gonyyi/afmt"
	"testing"
)

var VV bool
func vv(format string, a ...interface{}) {
	if VV {
		fmt.Printf(format, a...)
	}
}

func TestHumanNumber(t *testing.T) {
	type testPair struct {
		i int64
		s string
	}
	var expRes []testPair = []testPair{
		testPair{10, "10"},
		testPair{1000, "1.0K"},
		testPair{100000, "100.0K"},
		testPair{100000000, "100.0M"},
		testPair{100000000000, "100.0B"},
		testPair{100300000000, "100.3B"},
		testPair{100350000000, "100.3B"},
		testPair{100350000001, "100.4B"},
		testPair{10000000000000, "10.0T"},
	}

	for _, v := range expRes {
		res := afmt.HumanNumber(v.i, 1)
		vv("%7t: %d --> %s\n", v.s == res, v.i, res)
		if v.s != res {
			t.Errorf("unexpected result -- exp:%s, act:%s", v.s, res)
			t.Fail()
		}
	}
}

func TestNumberWithComma(t *testing.T) {
	type testPair struct {
		i int64
		s string
	}
	var expRes []testPair = []testPair{
		testPair{10, "10"},
		testPair{1000, "1,000"},
		testPair{100000, "100,000"},
		testPair{100000000, "100,000,000"},
		testPair{100000000000, "100,000,000,000"},
		testPair{10000000000000, "10,000,000,000,000"},
	}

	for _, v := range expRes {
		res := afmt.NumberWithComma(v.i)
		vv("%7t: %d --> %s\n", v.s == res, v.i, res)
		if v.s != res {
			t.Errorf("unexpected result -- exp:%s, act:%s", v.s, res)
			t.Fail()
		}
	}
}

func TestHumanBytes(t *testing.T) {
	expRes := map[int64]string{
		1:                "1B",
		1024:             "1.00KB",
		1048576:          "1.00MB",
		1073741824:       "1.00GB",
		1109941824:       "1.03GB",
		1099511627776:    "1.00TB",
		1125899906842624: "1.00PB",
	}

	sign := "=="
	vv("%8s | %8s | %8s\n", "Actual  ", "Expct'd ", "Invalid ")
	vv("%8s | %8s | %8s\n", "--------", "--------", "--------")
	for i, v := range expRes {
		hbyte := afmt.HumanBytes(i, 2)
		if hbyte != v {
			sign = "invalid"
			t.Errorf("unexpected result -- exp:%s, act:%s", v, hbyte)
			t.Fail()
		} else {
			sign = ""
		}
		vv("%8s | %8s | %8s\n", hbyte, v, sign)
	}
}
