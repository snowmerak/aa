package aa_test

import (
	"testing"

	"github.com/snowmerak/aa"
)

func TestMapAssertion_Equal(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).Equal(map[string]int{"a": 1, "b": 2})
	aa.AssertMap(t, map[string]int{}).Equal(map[string]int{})
}

func TestMapAssertion_NotEqual(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).NotEqual(map[string]int{"a": 1, "b": 3})
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).NotEqual(map[string]int{"a": 1})
}

func TestMapAssertion_ContainsAllKey(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).ContainsAllKey("a", "b")
}

func TestMapAssertion_NotContainsAllKey(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).NotContainsAllKey("c", "d")
}

func TestMapAssertion_ContainsAllValue(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).ContainsAllValue(1, 2)
}

func TestMapAssertion_NotContainsAllValue(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).NotContainsAllValue(3, 4)
}

func TestMapAssertion_ContainsAnyKey(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).ContainsAnyKey("a", "c")
}

func TestMapAssertion_NotContainsAnyKey(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).NotContainsAnyKey("c", "d")
}

func TestMapAssertion_ContainsAnyValue(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).ContainsAnyValue(1, 3)
}

func TestMapAssertion_NotContainsAnyValue(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).NotContainsAnyValue(3, 4)
}

func TestMapAssertion_UniqueValue(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).UniqueValue()
}

func TestMapAssertion_Empty(t *testing.T) {
	aa.AssertMap(t, map[string]int{}).Empty()
}

func TestMapAssertion_NotEmpty(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1}).NotEmpty()
}

func TestMapAssertion_Nil(t *testing.T) {
	aa.AssertMap(t, map[string]int(nil)).Nil()
}

func TestMapAssertion_NotNil(t *testing.T) {
	aa.AssertMap(t, map[string]int{}).NotNil()
}

func TestMapAssertion_LenEqual(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).LenEqual(2)
}

func TestMapAssertion_LenNotEqual(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).LenNotEqual(3)
}

func TestMapAssertion_LenGreaterThan(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).LenGreaterThan(1)
}

func TestMapAssertion_LenLessThan(t *testing.T) {
	aa.AssertMap(t, map[string]int{"a": 1, "b": 2}).LenLessThan(3)
}
