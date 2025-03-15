package aa_test

import (
	"testing"

	"github.com/snowmerak/aa"
)

func TestSliceEqual(t *testing.T) {
	if !aa.SliceEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}
}
