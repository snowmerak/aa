package aa

import (
	"errors"
	"strings"
	"testing"
)

type ErrorAssertion struct {
	t   *testing.T
	err error
}

func AssertError(t *testing.T, err error) *ErrorAssertion {
	return &ErrorAssertion{t: t, err: err}
}

func (ea *ErrorAssertion) IsNil() *ErrorAssertion {
	if ea.err != nil {
		ea.t.Fatalf("error is not nil")
	}
	return ea
}

func (ea *ErrorAssertion) IsNotNil() *ErrorAssertion {
	if ea.err == nil {
		ea.t.Fatalf("error is nil")
	}
	return ea
}

func (ea *ErrorAssertion) Equal(expected error) *ErrorAssertion {
	if ea.err != expected {
		ea.t.Fatalf("error not equal: input[%v] != expectec[%v]", ea.err, expected)
	}
	return ea
}

func (ea *ErrorAssertion) NotEqual(expected error) *ErrorAssertion {
	if ea.err == expected {
		ea.t.Fatalf("error equal: input[%v] == expectec[%v]", ea.err, expected)
	}
	return ea
}

func (ea *ErrorAssertion) Is(expected error) *ErrorAssertion {
	if !errors.Is(ea.err, expected) {
		ea.t.Fatalf("error contains: input[%v] != expectec[%v]", ea.err, expected)
	}
	return ea
}

func (ea *ErrorAssertion) IsNot(expected error) *ErrorAssertion {
	if errors.Is(ea.err, expected) {
		ea.t.Fatalf("error not contains: input[%v] == expectec[%v]", ea.err, expected)
	}
	return ea
}

func (ea *ErrorAssertion) MessageContains(expected string) *ErrorAssertion {
	if !strings.Contains(ea.err.Error(), expected) {
		ea.t.Fatalf("error not contains message: expected[%s] in input[%v]", expected, ea.err)
	}
	return ea
}

func (ea *ErrorAssertion) MessageNotContains(expected string) *ErrorAssertion {
	if strings.Contains(ea.err.Error(), expected) {
		ea.t.Fatalf("error contains message: expected[%s] in input[%v]", expected, ea.err)
	}
	return ea
}

func (ea *ErrorAssertion) MessageEqual(expected string) *ErrorAssertion {
	if ea.err.Error() != expected {
		ea.t.Fatalf("error message not equal: input[%v] != expectec[%v]", ea.err, expected)
	}
	return ea
}

func (ea *ErrorAssertion) MessageNotEqual(expected string) *ErrorAssertion {
	if ea.err.Error() == expected {
		ea.t.Fatalf("error message equal: input[%v] == expectec[%v]", ea.err, expected)
	}
	return ea
}
