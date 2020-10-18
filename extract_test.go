package afmt_test

import (
	"github.com/gonyyi/afmt"
	"testing"
)

func TestNewExtract(t *testing.T) {
	s := "https://blog.gonyyi.com/myLife/123/"
	// ============================== TEST FOR 2 VARS
	{
		ext := afmt.NewExtract("^https://blog.gonyyi.com/([^/]+)/([^/]+)/?$", 2)(s)
		if ext[0] != "myLife" || ext[1] != "123" {
			t.Errorf("Expected: [myLife, 123], Received: [%s, %s]", ext[0], ext[1])
		}
	}

	// ============================== TEST FOR 1 VAR
	{
		ext := afmt.NewExtract("^https://blog.gonyyi.com/([^/]+)/[^/]+/?$", 1)(s)
		if ext[0] != "myLife" {
			t.Errorf("Expected: [myLife], Received: [%s]", ext[0])
		}
	}

	// ============================== TEST FOR STANDALONE EXTRACT
	{
		ext := afmt.Extract(s, "^https://blog.gonyyi.com/([^/]+)/([^/]+)/?$", 2)
		if ext[0] != "myLife" || ext[1] != "123" {
			t.Errorf("Expected: [myLife, 123], Received: [%s, %s]", ext[0], ext[1])
		}
	}
}

func BenchmarkExtract(b *testing.B) {
	// BenchmarkExtract-12    	 2108232	       565 ns/op	     112 B/op	       2 allocs/op

	extFn := afmt.NewExtract("^https://blog.gonyyi.com/([^/]+)/([^/]+)/?$", 2)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		extFn("https://blog.gonyyi.com/myLife/123/")
	}
}
