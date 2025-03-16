package aa_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/snowmerak/aa"
)

func TestErrorAssertion_IsNil(t *testing.T) {
	var err error
	aa.AssertError(t, err).IsNil()
}

func TestErrorAssertion_IsNotNil(t *testing.T) {
	err := errors.New("test error")
	aa.AssertError(t, err).IsNotNil()
}

func TestErrorAssertion_Equal(t *testing.T) {
	err := errors.New("test error")
	aa.AssertError(t, err).Equal(err)
}

func TestErrorAssertion_NotEqual(t *testing.T) {
	err1 := errors.New("test error")
	err2 := errors.New("test error")
	aa.AssertError(t, err1).NotEqual(err2)
}

func TestErrorAssertion_Is(t *testing.T) {
	baseErr := errors.New("base error")
	wrappedErr := fmt.Errorf("wrapped: %w", baseErr)
	aa.AssertError(t, wrappedErr).Is(baseErr)
}

func TestErrorAssertion_IsNot(t *testing.T) {
	baseErr := errors.New("base error")
	wrappedErr := fmt.Errorf("wrapped: %w", baseErr)
	aa.AssertError(t, wrappedErr).IsNot(errors.New("another error"))
}

func TestErrorAssertion_MessageContains(t *testing.T) {
	err := errors.New("test error message")
	aa.AssertError(t, err).MessageContains("error")
}

func TestErrorAssertion_MessageNotContains(t *testing.T) {
	err := errors.New("test error message")
	aa.AssertError(t, err).MessageNotContains("another")
}

func TestErrorAssertion_MessageEqual(t *testing.T) {
	err := errors.New("test error")
	aa.AssertError(t, err).MessageEqual("test error")
}

func TestErrorAssertion_MessageNotEqual(t *testing.T) {
	err := errors.New("test error")
	aa.AssertError(t, err).MessageNotEqual("another error")
}
