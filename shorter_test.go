package afmt_test

import (
	"github.com/gonyyi/afmt"
	"testing"
)

func TestLeftAndRight(t *testing.T) {
	// VV = true

	// Left
	vv("\nNewLeft() ------\n")
	{
		f := afmt.NewLeft(10)
		e := func(inp, expRes string) {
			tmp := f(inp)
			var isGood string = "OK"
			if tmp != expRes {isGood = "FAIL"}

			vv("%-4s  [%s] --> [%s] (len: %d)\n", isGood, inp, tmp, len(tmp))
			if tmp != expRes {
				t.Fail()
			}
		}
		e("hello",              "hello     ")
		e("gon",                "gon       ")
		e("gon is the coolest", "gon is the")
	}

	vv("\nNewRight() ------\n")
	// Right
	{
		f := afmt.NewRight(10)
		e := func(inp, expRes string) {
			tmp := f(inp)
			var isGood string = "OK"
			if tmp != expRes {isGood = "FAIL"}
			vv("%-4s  [%s] --> [%s] (len: %d)\n", isGood, inp, tmp, len(tmp))
			if tmp != expRes {
				t.Fail()
			}
		}
		e("hello",              "hello     ")
		e("gon",                "gon       ")
		e("gon is the coolest", "he coolest")
	}


}

func TestNewShorterFunc(t *testing.T) {
	const testString = "123456789ABCDEF"
	const testMarker = ".."
	t.Run("common cases", func(t2 *testing.T) {
		if v := afmt.NewShorterFunc(9, testMarker, 4, 2)(testString); v != "12345..EF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345..EF") // Exp: 12345..EF
		}
		if v := afmt.NewShorterFunc(9, testMarker, 2, 4)(testString); v != "12..BCDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12..BCDEF") // Exp: 12..BCDEF
		}
		if v := afmt.NewShorterFunc(9, testMarker, 4, 4)(testString); v != "1234..DEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "1234..DEF") // Exp: 1234..DEF
		}
		if v := afmt.NewShorterFunc(10, testMarker, 4, 4)(testString); v != "1234..CDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "1234..CDEF") // Exp: 1234..CDEF
		}
		if v := afmt.NewShorterFunc(10, testMarker, 0, 0)(testString); v != "1234..CDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "1234..CDEF") // Exp: 1234..CDEF
		}
		if v := afmt.NewShorterFunc(10, testMarker, 0, 1)(testString); v != "..89ABCDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "..89ABCDEF") // Exp: ..89ABCDEF
		}
		if v := afmt.NewShorterFunc(10, testMarker, 1, 0)(testString); v != "12345678.." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345678..") // Exp: 12345678..
		}
		if v := afmt.NewShorterFunc(10, testMarker, 20, 0)(testString); v != "12345678.." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345678..") // Exp: 12345678..
		}
		if v := afmt.NewShorterFunc(10, testMarker, 0, 20)(testString); v != "..89ABCDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "..89ABCDEF") // Exp: ..89ABCDEF
		}
		if v := afmt.NewShorterFunc(10, testMarker, 3, 20)(testString); v != "123..BCDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "123..BCDEF") // Exp: 123..BCDEF })
		}
	})

	t.Run("weird cases", func(t2 *testing.T) {
		if v := afmt.NewShorterFunc(10, testMarker, 20, 20)(testString); v != "1234..CDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "1234..CDEF") // Exp: 1234..CDEF
		}
		if v := afmt.NewShorterFunc(10, testMarker, 25, 20)(testString); v != "..89ABCDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "..89ABCDEF") // Exp: ..89ABCDEF
		}
		if v := afmt.NewShorterFunc(10, testMarker, 25, 29)(testString); v != "12345678.." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345678..") // Exp: 12345678..
		}
		if v := afmt.NewShorterFunc(10, testMarker, 25, 5)(testString); v != "123..BCDEF" {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "123..BCDEF") // Exp:  123..BCDEF
		}
		if v := afmt.NewShorterFunc(2, testMarker, 25, 5)(testString); v != ".." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "..") // Exp:  ..
		}
		if v := afmt.NewShorterFunc(2, "/.-", 25, 5)(testString); v != "/." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "/.") // Exp:  /.
		}
		if v := afmt.NewShorterFunc(-10, testMarker, 25, 29)(testString); v != "12345678.." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345678..") // Exp: 12345678..
		}
		if v := afmt.NewShorterFunc(10, testMarker, -25, 29)(testString); v != "12345678.." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345678..") // Exp: 12345678..
		}
		if v := afmt.NewShorterFunc(10, testMarker, 25, -29)(testString); v != "12345678.." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345678..") // Exp: 12345678..
		}
		if v := afmt.NewShorterFunc(-10, testMarker, -25, -29)(testString); v != "12345678.." {
			t2.Errorf("unexpected - rec: %s, exp: %s", v, "12345678..") // Exp: 12345678..
		}
	})
}
