package aa

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type NumberAssertion[T Number] struct {
	t *testing.T
	n T
}

func AssertNumber[T Number](t *testing.T, n T) *NumberAssertion[T] {
	return &NumberAssertion[T]{t: t, n: n}
}

func (na *NumberAssertion[T]) Equal(b T) *NumberAssertion[T] {
	if na.n != b {
		na.t.Fatalf("numbers are not equal: a=%v, b=%v", na.n, b)
	}

	return na
}

func (na *NumberAssertion[T]) NotEqual(b T) *NumberAssertion[T] {
	if na.n == b {
		na.t.Fatalf("numbers are equal: a=%v, b=%v", na.n, b)
	}

	return na
}

func (na *NumberAssertion[T]) GreaterThan(b T) *NumberAssertion[T] {
	if na.n <= b {
		na.t.Fatalf("number is not greater than: a=%v, b=%v", na.n, b)
	}

	return na
}

func (na *NumberAssertion[T]) LessThan(b T) *NumberAssertion[T] {
	if na.n >= b {
		na.t.Fatalf("number is not less than: a=%v, b=%v", na.n, b)
	}

	return na
}

func (na *NumberAssertion[T]) Zero() *NumberAssertion[T] {
	if !reflect.ValueOf(na.n).IsZero() {
		na.t.Fatalf("number is not zero: a=%v", na.n)
	}

	return na
}

func (na *NumberAssertion[T]) NotZero() *NumberAssertion[T] {
	if reflect.ValueOf(na.n).IsZero() {
		na.t.Fatalf("number is zero: a=%v", na.n)
	}

	return na
}

func (na *NumberAssertion[T]) Positive() *NumberAssertion[T] {
	if na.n < 0 {
		na.t.Fatalf("number is not positive: a=%v", na.n)
	}

	return na
}

func (na *NumberAssertion[T]) Negative() *NumberAssertion[T] {
	if na.n > 0 {
		na.t.Fatalf("number is not negative: a=%v", na.n)
	}

	return na
}

func (na *NumberAssertion[T]) Even() *NumberAssertion[T] {
	a := any(na.n)
	b := false
	switch v := a.(type) {
	case int8:
		if v&1 == 0 {
			b = true
		}
	case int16:
		if v&1 == 0 {
			b = true
		}
	case int32:
		if v&1 == 0 {
			b = true
		}
	case int64:
		if v&1 == 0 {
			b = true
		}
	case int:
		if v&1 == 0 {
			b = true
		}
	case uint8:
		if v&1 == 0 {
			b = true
		}
	case uint16:
		if v&1 == 0 {
			b = true
		}
	case uint32:
		if v&1 == 0 {
			b = true
		}
	case uint64:
		if v&1 == 0 {
			b = true
		}
	case uint:
		if v&1 == 0 {
			b = true
		}
	case float64:
		if int64(v)%2 == 0 {
			b = true
		}
	case float32:
		if int32(v)%2 == 0 {
			b = true
		}
	default:
		na.t.Fatalf("Even: unsupported type: %T", a)
	}

	if !b {
		na.t.Fatalf("number is not even: a=%v", na.n)
	}

	return na
}

func (na *NumberAssertion[T]) Odd() *NumberAssertion[T] {
	a := any(na.n)
	b := false
	switch v := a.(type) {
	case int8:
		if v&1 == 1 {
			b = true
		}
	case int16:
		if v&1 == 1 {
			b = true
		}
	case int32:
		if v&1 == 1 {
			b = true
		}
	case int64:
		if v&1 == 1 {
			b = true
		}
	case int:
		if v&1 == 1 {
			b = true
		}
	case uint8:
		if v&1 == 1 {
			b = true
		}
	case uint16:
		if v&1 == 1 {
			b = true
		}
	case uint32:
		if v&1 == 1 {
			b = true
		}
	case uint64:
		if v&1 == 1 {
			b = true
		}
	case uint:
		if v&1 == 1 {
			b = true
		}
	case float64:
		if int64(v)%2 == 1 {
			b = true
		}
	case float32:
		if int32(v)%2 == 1 {
			b = true
		}
	default:
		na.t.Fatalf("Odd: unsupported type: %T", a)
	}

	if !b {
		na.t.Fatalf("number is not odd: a=%v", na.n)
	}

	return na
}
