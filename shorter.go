package afmt

import (
	"strings"
)

// =====================================================================================================================
// LEFT
// =====================================================================================================================

func NewLeft(n int) func(string) string {
	return func(s string) string {
		return Left(s, n)
	}
}

func Left(s string, n int) string {
	lenS := len(s)
	if lenS >= n {
		return s[0:n]
	}
	return s + strings.Repeat(" ", n-lenS)
}

// =====================================================================================================================
// RIGHT
// =====================================================================================================================

func NewRight(n int) func(string) string {
	return func(s string) string {
		return Right(s, n)
	}
}

func Right(s string, n int) string {
	lenS := len(s)
	if lenS >= n {
		return s[len(s)-n:]
	}
	return s + strings.Repeat(" ", n-lenS)
}


// NewShorterFunc will make string at target output length.
func NewShorterFunc(tgtOutLength int, dotdotMarker string, minLeft, minRight int) func(string) string {
	// =================================================================================================================
	// USAGE:
	//   dfmt.NewShortyFunc(7, "...", 0, 0)("R8O5DpWJbDbzKdHpg82HqQAtXz68BIxOV8RqlhEB3") // returns "R8...B3"
	//   dfmt.NewShortyFunc(7, "...", 5, 0)("R8O5DpWJbDbzKdHpg82HqQAtXz68BIxOV8RqlhEB3") // returns "R8...B3"
	//   dfmt.NewShortyFunc(7, "...", 0, 5)("R8O5DpWJbDbzKdHpg82HqQAtXz68BIxOV8RqlhEB3") // returns "R8...B3"
	//   dfmt.NewShortyFunc(7, "...", 2, 5)("R8O5DpWJbDbzKdHpg82HqQAtXz68BIxOV8RqlhEB3") // returns "R8...B3"
	//      When minLeft/minRight are equal, evenly
	//      Else whoever smallest sets the actual size (largest will be larger than min)
	// =================================================================================================================
	ABS := func(i int) int {
		if i < 0 {
			i = i * -1
		}
		return i
	}

	// Only take positive number and just make it positive...
	tgtOutLength = ABS(tgtOutLength)
	minLeft = ABS(minLeft)
	minRight = ABS(minRight)

	var left, right int
	intDotdotMarker := len(dotdotMarker)

	if tgtOutLength == intDotdotMarker {
		return func(string) string {
			return dotdotMarker
		}
	}
	if tgtOutLength < intDotdotMarker {
		return func(string) string {
			return dotdotMarker[0:tgtOutLength]
		}
	}

	if minLeft == minRight { // equally
		if tgtOutLength < (intDotdotMarker + 2) {
			tgtOutLength = intDotdotMarker + 2
		}
		right = (tgtOutLength - intDotdotMarker) / 2
		left = (tgtOutLength - intDotdotMarker) - right
	} else if minLeft > minRight { // right length fixed
		// 123456789ABCDEF / targetLength=9, market=".."
		// assume (4, 2) --> 12345..EF
		left = (tgtOutLength - intDotdotMarker) - minRight // 9-2-2 = 5
		right = (tgtOutLength - intDotdotMarker) - left    // 9-2-5 = 2

	} else if minLeft < minRight { // left length fixed
		// 123456789ABCDEF / targetLength=9, market=".."
		// assume (2, 4) --> 12..BCDEF
		right = (tgtOutLength - intDotdotMarker) - minLeft // 9 - 2 - 2 = 5
		left = (tgtOutLength - intDotdotMarker) - right    // 9 - 2 - 5 = 2
	}

	if tgtOutLength < left || tgtOutLength < right {
		if left > right { // show only left - dotdot marker
			return func(s string) string {
				return s[0:(tgtOutLength-intDotdotMarker)] + dotdotMarker
			}
		} else if left < right {
			return func(s string) string {
				return dotdotMarker + s[len(s)-tgtOutLength+intDotdotMarker:]
			}
		}
	}

	return func(s string) string {
		if len(s) <= tgtOutLength {
			return s
		}
		return s[0:left] + dotdotMarker + s[len(s)-right:]
	}
}

