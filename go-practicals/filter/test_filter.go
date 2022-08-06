package filter

import (
	"testing"
)

//NOTE: this test is broken - need to investigate testing that compares slices of values
func TestFilter(t *testing.T) {
	if filter(isOdd, [4]int{1, 2, 3, 4}) != [2]int{1, 3} {

		t.Error("filter function isn't filtering")
	}
}
