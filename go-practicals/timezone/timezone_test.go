package timezone

import (
	"fmt"
	"testing"
)

// improper testing: no assertions made
func TestTzConvert(t *testing.T) {
	ts := "1999-05-10T10:00"
	out, err := tzConvert(ts, "America/Los_Angeles", "Asia/Jerusalem")
	if err != nil {
		fmt.Errorf("error: %q", err)
		return
	}
	fmt.Println(out)

}
