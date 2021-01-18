package afmt

import (
  "fmt"
  "strconv"
)

func BitString(flag uint64) string {
	return fmt.Sprintf("%064s", strconv.FormatUint(flag, 2))
}
