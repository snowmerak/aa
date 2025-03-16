package aa

import "testing"

type BoolAssertion struct {
	t *testing.T
	b bool
}

func AssertBool(t *testing.T, b bool) *BoolAssertion {
	return &BoolAssertion{t: t, b: b}
}

func (ba *BoolAssertion) Equal(expected bool) *BoolAssertion {
	if ba.b != expected {
		ba.t.Fatalf("expected %v, got %v", expected, ba.b)
	}
	return ba
}

func (ba *BoolAssertion) NotEqual(expected bool) *BoolAssertion {
	if ba.b == expected {
		ba.t.Fatalf("expected %v, got %v", !expected, ba.b)
	}
	return ba
}

func (ba *BoolAssertion) IsTrue() *BoolAssertion {
	if !ba.b {
		ba.t.Fatalf("expected true, got false")
	}
	return ba
}

func (ba *BoolAssertion) IsFalse() *BoolAssertion {
	if ba.b {
		ba.t.Fatalf("expected false, got true")
	}
	return ba
}
