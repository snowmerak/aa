package aa

import (
	"fmt"
	"reflect"
	"testing"
)

type InterfaceAssertion struct {
	t *testing.T
	a any
}

func AssertInterface(t *testing.T, a any) *InterfaceAssertion {
	return &InterfaceAssertion{t: t, a: a}
}

func (ia *InterfaceAssertion) Equal(b any) *InterfaceAssertion {
	if ia.a != b {
		ia.t.Fatalf("interface not equal: input[%v] != expected[%v]", ia.a, b)
	}
	return ia
}

func (ia *InterfaceAssertion) NotEqual(b any) *InterfaceAssertion {
	if ia.a == b {
		ia.t.Fatalf("interface equal: input[%v] == expected[%v]", ia.a, b)
	}
	return ia
}

func (ia *InterfaceAssertion) IsNil() *InterfaceAssertion {
	if ia.a != nil {
		ia.t.Fatalf("interface not nil: input[%v]", ia.a)
	}
	return ia
}

func (ia *InterfaceAssertion) IsNotNil() *InterfaceAssertion {
	if ia.a == nil {
		ia.t.Fatalf("interface nil: input[%v]", ia.a)
	}
	return ia
}

func (ia *InterfaceAssertion) DeepEqual(b any) *InterfaceAssertion {
	if !reflect.ValueOf(ia.a).Equal(reflect.ValueOf(b)) {
		ia.t.Fatalf("interface not deep equal: input[%v] != expected[%v]", ia.a, b)
	}
	return ia
}

func (ia *InterfaceAssertion) NotDeepEqual(b any) *InterfaceAssertion {
	if reflect.ValueOf(ia.a).Equal(reflect.ValueOf(b)) {
		ia.t.Fatalf("interface deep equal: input[%v] == expected[%v]", ia.a, b)
	}
	return ia
}

func (ia *InterfaceAssertion) IsTypeOf(b any) *InterfaceAssertion {
	if reflect.TypeOf(ia.a) != reflect.TypeOf(b) {
		ia.t.Fatalf("interface not type of: input[%T] != expected[%T]", ia.a, b)
	}
	return ia
}

func (ia *InterfaceAssertion) IsNotTypeOf(b any) *InterfaceAssertion {
	if reflect.TypeOf(ia.a) == reflect.TypeOf(b) {
		ia.t.Fatalf("interface type of: input[%T] == expected[%T]", ia.a, b)
	}
	return ia
}

// Deprecated: do not use this
func (ia *InterfaceAssertion) Implements(b any) *InterfaceAssertion {
	fmt.Printf("%v\n", reflect.TypeOf(b))
	if reflect.TypeOf(ia.a).Implements(reflect.TypeOf(b)) {
		ia.t.Fatalf("interface not implements: input[%T] in expected[%T]", ia.a, b)
	}
	return ia
}

// Deprecated: do not use this
func (ia *InterfaceAssertion) NotImplements(b any) *InterfaceAssertion {
	if !reflect.TypeOf(ia.a).Implements(reflect.TypeOf(b)) {
		ia.t.Fatalf("interface implements: input[%T] in expected[%T]", ia.a, b)
	}
	return ia
}
