package afmt_test

import (
	"fmt"
	"github.com/gonyyi/afmt"
	"testing"
)

func TestHyphenate(t *testing.T) {
	// VV = true // this enables vv() to be printed

	type result struct {
		s   string
		sep string
		spc []int
		exp string
	}

	expRes := []result{
		result{"abcdefgh", "-", []int{1, 3, 4}, "a-bcd-efgh"},
		result{"abcdefgh", ".", []int{2, 2, 2, 2}, "ab.cd.ef.gh"},
		result{"012345678", "-", []int{3, 2, 4}, "012-34-5678"},
		result{"1234567890123456", "-", []int{4, 4, 4, 4}, "1234-5678-9012-3456"},
		result{"123abc456def", " ", []int{3,3,3,3}, "123 abc 456 def"},
	}

	// anyFail := false
	for _, v := range expRes {
		out, ok := afmt.Hyphenate(v.s, v.sep, v.spc...)
		additional := ""
		if out != v.exp { additional = fmt.Sprintf("(exp: %s)", v.exp) }
		vv("%-30s --> %s %s\n", v.s, out, additional)

		if !ok {
			t.Errorf("failed to hyphenate [%s], [%s]", v.s, v.sep)
			t.Fail()
		}
	}

}
