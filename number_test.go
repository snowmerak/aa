package aa_test

import (
	"testing"

	"github.com/snowmerak/aa"
)

func TestNumberAssertion_Equal(t *testing.T) {
	aa.AssertNumber(t, 5).Equal(5)
	aa.AssertNumber(t, 5.5).Equal(5.5)
}

func TestNumberAssertion_NotEqual(t *testing.T) {
	aa.AssertNumber(t, 5).NotEqual(6)
	aa.AssertNumber(t, 5.5).NotEqual(6.5)
}

func TestNumberAssertion_GreaterThan(t *testing.T) {
	aa.AssertNumber(t, 5).GreaterThan(4)
	aa.AssertNumber(t, 5.5).GreaterThan(4.5)
}

func TestNumberAssertion_LessThan(t *testing.T) {
	aa.AssertNumber(t, 5).LessThan(6)
	aa.AssertNumber(t, 5.5).LessThan(6.5)
}

func TestNumberAssertion_Zero(t *testing.T) {
	aa.AssertNumber(t, 0).Zero()
	aa.AssertNumber(t, 0.0).Zero()
}

func TestNumberAssertion_NotZero(t *testing.T) {
	aa.AssertNumber(t, 5).NotZero()
	aa.AssertNumber(t, 5.5).NotZero()
}

func TestNumberAssertion_Positive(t *testing.T) {
	aa.AssertNumber(t, 5).Positive()
	aa.AssertNumber(t, 5.5).Positive()
}

func TestNumberAssertion_Negative(t *testing.T) {
	aa.AssertNumber(t, -5).Negative()
	aa.AssertNumber(t, -5.5).Negative()
}

func TestNumberAssertion_Even(t *testing.T) {
	aa.AssertNumber(t, 4).Even()
	aa.AssertNumber(t, 4.0).Even()
}

func TestNumberAssertion_Odd(t *testing.T) {
	aa.AssertNumber(t, 5).Odd()
	aa.AssertNumber(t, 5.0).Odd()
}
