package aa_test

import (
	"testing"

	"github.com/snowmerak/aa"
)

func TestBoolAssertion_Equal(t *testing.T) {
	aa.AssertBool(t, true).Equal(true)
	aa.AssertBool(t, false).Equal(false)
}

func TestBoolAssertion_NotEqual(t *testing.T) {
	aa.AssertBool(t, true).NotEqual(false)
	aa.AssertBool(t, false).NotEqual(true)
}

func TestBoolAssertion_IsTrue(t *testing.T) {
	aa.AssertBool(t, true).IsTrue()
}

func TestBoolAssertion_IsFalse(t *testing.T) {
	aa.AssertBool(t, false).IsFalse()
}
