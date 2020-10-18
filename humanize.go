package afmt

import (
	"fmt"
	"regexp"
)

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

const (
	num_k = 1000
	num_m = num_k * 1000
	num_b = num_m * 1000
	num_t = num_b * 1000
)

func NumberWithComma(num int64) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}

func HumanNumber(i int64, decimal uint) string {
	isNegative := ""
	if i < 0 {
		isNegative = "-"
		i = i * -1
	}

	var newInt float64 // because of decimal point
	unit := ""

	if i >= num_t {
		unit = "T"
		newInt = float64(i) / num_t
	} else if i >= num_b {
		unit = "B"
		newInt = float64(i) / num_b
	} else if i >= num_m {
		unit = "M"
		newInt = float64(i) / num_m
	} else if i >= num_k {
		unit = "K"
		newInt = float64(i) / num_k
	} else { // below 1k
		return fmt.Sprintf("%d", i)
	}

	ftr := fmt.Sprintf("%%s%%.%df%s", decimal, unit)
	return fmt.Sprintf(ftr, isNegative, newInt)
}

// HumanBytes takes int64 and convers to human readable format such as
// 1073741824 --> 1.0GB
// The reason this only take int64 is other int aren't as accurate.
func HumanBytes(i int64, decimal uint) string {
	isNegative := ""
	if i < 0 {
		isNegative = "-"
		i = i * -1
	}
	i2 := float64(i)

	var newInt float64 // because of decimal point
	unit := ""
	if i2 >= PB {
		unit = "PB"
		newInt = i2 / PB
	} else if i2 >= TB {
		unit = "TB"
		newInt = i2 / TB
	} else if i2 >= GB {
		unit = "GB"
		newInt = i2 / GB
	} else if i2 >= MB {
		unit = "MB"
		newInt = i2 / MB
	} else if i2 >= KB {
		unit = "KB"
		newInt = i2 / KB
	} else { // below 1k
		return fmt.Sprintf("%dB", i) // i1 is int64 type
	}

	ftr := fmt.Sprintf("%%s%%.%df%s", decimal, unit)
	return fmt.Sprintf(ftr, isNegative, newInt)
}
