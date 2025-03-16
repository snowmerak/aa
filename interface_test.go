package aa_test

import (
	"testing"

	"github.com/snowmerak/aa"
)

func TestInterfaceAssertion_Equal(t *testing.T) {
	ia := aa.AssertInterface(t, "test")
	ia.Equal("test")
}

func TestInterfaceAssertion_NotEqual(t *testing.T) {
	ia := aa.AssertInterface(t, "test")
	ia.NotEqual("different")
}

func TestInterfaceAssertion_IsNil(t *testing.T) {
	ia := aa.AssertInterface(t, nil)
	ia.IsNil()
}

func TestInterfaceAssertion_IsNotNil(t *testing.T) {
	ia := aa.AssertInterface(t, "test")
	ia.IsNotNil()
}

func TestInterfaceAssertion_DeepEqual(t *testing.T) {
	ia := aa.AssertInterface(t, [3]int{1, 2, 3})
	ia.DeepEqual([3]int{1, 2, 3})
}

func TestInterfaceAssertion_NotDeepEqual(t *testing.T) {
	ia := aa.AssertInterface(t, [3]int{1, 2, 3})
	ia.NotDeepEqual([3]int{4, 5, 6})
}

func TestInterfaceAssertion_IsTypeOf(t *testing.T) {
	ia := aa.AssertInterface(t, "test")
	ia.IsTypeOf("another test")
}

func TestInterfaceAssertion_IsNotTypeOf(t *testing.T) {
	ia := aa.AssertInterface(t, "test")
	ia.IsNotTypeOf(123)
}

// func TestInterfaceAssertion_Implements(t *testing.T) {
//	ia := aa.AssertInterface(t, "test")
//	ia.Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem())
// }
//
// func TestInterfaceAssertion_NotImplements(t *testing.T) {
//	ia := aa.AssertInterface(t, "test")
//	ia.NotImplements((testing.TB)(nil))
// }
