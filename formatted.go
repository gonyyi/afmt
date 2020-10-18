package afmt

import "bytes"

// Hyphenate expects the output length to be [totalNumberOfDelimiter + string].
// If the length does not match, it will return false to "ok bool".
func Hyphenate(s, delimiter string, spaceBetweenSep ...int) (out string, ok bool) {
	// if func("abcdefghi", "-", 3, 2, 4) ==> "abc-de-fghi"
	bs := []byte(s)
	dlmtr := []byte(delimiter)
	lenLength := len(spaceBetweenSep)
	var starts = make([]int, len(spaceBetweenSep))
	sIdx := 0 // idex for start
	totLen := 0
	for idx, v := range spaceBetweenSep {
		totLen += v
		starts[idx] = sIdx
		sIdx += v
	}

	// ============================== CHECK IF LENGTH MATCHES
	if len(s) != totLen {
		return "", false
	}
	var buf bytes.Buffer
	for idx, v := range starts {
		// println(idx, v, spaceBetweenSep[idx], string(bs[v:v+spaceBetweenSep[idx]]))
		buf.Write(bs[v : v+spaceBetweenSep[idx]])
		if idx < lenLength-1 {
			buf.Write(dlmtr)
		}
	}
	return buf.String(), true
}
