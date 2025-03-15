package aa_test

import (
	"testing"

	"github.com/snowmerak/aa"
)

func TestSliceAssertion_Equal(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).Equal([]int{1, 2, 3})
	aa.AssertSlice(t, []int{}).Equal([]int{})
}

func TestSliceAssertion_NotEqual(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).NotEqual([]int{1, 2, 4})
	aa.AssertSlice(t, []int{1, 2, 3}).NotEqual([]int{1, 2})
}

func TestSliceAssertion_Contains(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).Contains(2)
}

func TestSliceAssertion_NotContains(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).NotContains(4)
}

func TestSliceAssertion_ContainsAll(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).ContainsAll(1, 2)
}

func TestSliceAssertion_NotContainsAll(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).NotContainsAll(4, 5)
}

func TestSliceAssertion_ContainsAny(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).ContainsAny(2, 4)
}

func TestSliceAssertion_NotContainsAny(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).NotContainsAny(4, 5)
}

func TestSliceAssertion_Unique(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).Unique()
}

func TestSliceAssertion_Same(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).Same([]int{1, 2, 3})
}

func TestSliceAssertion_Empty(t *testing.T) {
	aa.AssertSlice(t, []int{}).Empty()
}

func TestSliceAssertion_Nil(t *testing.T) {
	aa.AssertSlice(t, []int(nil)).Nil()
}

func TestSliceAssertion_NotEmpty(t *testing.T) {
	aa.AssertSlice(t, []int{1}).NotEmpty()
}

func TestSliceAssertion_NotNil(t *testing.T) {
	aa.AssertSlice(t, []int{}).NotNil()
}

func TestSliceAssertion_LenEqual(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).LenEqual(3)
}

func TestSliceAssertion_LenNotEqual(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).LenNotEqual(2)
}

func TestSliceAssertion_LenGreaterThan(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).LenGreaterThan(2)
}

func TestSliceAssertion_LenLessThan(t *testing.T) {
	aa.AssertSlice(t, []int{1, 2, 3}).LenLessThan(4)
}

func TestSliceAssertion_CapEqual(t *testing.T) {
	aa.AssertSlice(t, make([]int, 3, 5)).CapEqual(5)
}

func TestSliceAssertion_CapNotEqual(t *testing.T) {
	aa.AssertSlice(t, make([]int, 3, 5)).CapNotEqual(4)
}

func TestSliceAssertion_CapGreaterThan(t *testing.T) {
	aa.AssertSlice(t, make([]int, 3, 5)).CapGreaterThan(4)
}

func TestSliceAssertion_CapLessThan(t *testing.T) {
	aa.AssertSlice(t, make([]int, 3, 5)).CapLessThan(6)
}
