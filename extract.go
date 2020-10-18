package afmt

import (
	"regexp"
)

// NewExtract(string, int) func(string)[]string
// This will return a function that spit out a condition where regex matches.
// If none matches, it will return blank space to avoid null pointer error
// instead of skipping it.
// (note: using sync.Pool didn't have any allocation gain)
func NewExtract(fmtrRegex string, numFlds int) func(string) []string {
	if re, _ := regexp.Compile(fmtrRegex); re != nil {
		return func(s string) []string {
			tmp := re.FindStringSubmatch(s)
			if len(tmp) != (numFlds + 1) { // because first item in the array is whole string, it will have +1
				tmp = make([]string, numFlds)
				return tmp
			}
			return tmp[1:] // because first item in the array is whole string
		}
	}
	return func(string) []string {
		return make([]string, numFlds)
	}
}

// Extract() is a standalone version of NewExtract.
// Suitable to use when no repeat is expected.
func Extract(s, fmtrRegex string, numFlds int) []string {
	tmp := regexp.MustCompile(fmtrRegex).FindStringSubmatch(s)
	if len(tmp) != (numFlds + 1) { // because first item in the array is whole string, it will have +1
		tmp = make([]string, numFlds)
		return tmp
	}
	return tmp[1:] // because first item in the array is whole string
}
