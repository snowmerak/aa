package aa_test

import (
	"testing"

	"github.com/snowmerak/aa"
)

func TestBytesAssertion_Equal(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.Equal([]byte("test"))
}

func TestBytesAssertion_NotEqual(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.NotEqual([]byte("different"))
}

func TestBytesAssertion_IsEqualFold(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("TEST"))
	ba.IsEqualFold([]byte("test"))
}

func TestBytesAssertion_IsNotEqualFold(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("TEST"))
	ba.IsNotEqualFold([]byte("different"))
}

func TestBytesAssertion_Contains(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.Contains([]byte("es"))
}

func TestBytesAssertion_NotContains(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.NotContains([]byte("xyz"))
}

func TestBytesAssertion_ContainsAll(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.ContainsAll([]byte("te"), []byte("st"))
}

func TestBytesAssertion_NotContainsAll(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.NotContainsAll([]byte("te"), []byte("xyz"))
}

func TestBytesAssertion_ContainsAny(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.ContainsAny([]byte("te"), []byte("xyz"))
}

func TestBytesAssertion_NotContainsAny(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.NotContainsAny([]byte("xyz"), []byte("abc"))
}

func TestBytesAssertion_IsEmpty(t *testing.T) {
	ba := aa.AssertBytes(t, []byte(""))
	ba.IsEmpty()
}

func TestBytesAssertion_IsNotEmpty(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.IsNotEmpty()
}

func TestBytesAssertion_IsNil(t *testing.T) {
	ba := aa.AssertBytes(t, []byte(nil))
	ba.IsNil()
}

func TestBytesAssertion_IsNotNil(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.IsNotNil()
}

func TestBytesAssertion_HasPrefix(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.HasPrefix([]byte("te"))
}

func TestBytesAssertion_HasSuffix(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.HasSuffix([]byte("st"))
}

func TestBytesAssertion_LenEqual(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.LenEqual(4)
}

func TestBytesAssertion_LenNotEqual(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.LenNotEqual(5)
}

func TestBytesAssertion_LenGreaterThan(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.LenGreaterThan(3)
}

func TestBytesAssertion_LenLessThan(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.LenLessThan(5)
}

func TestBytesAssertion_EqualHex(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("test"))
	ba.EqualHex("74657374")
}

func TestBytesAssertion_EqualJSON(t *testing.T) {
	ba := aa.AssertBytes(t, []byte("{\"key\":\"value\"}\n"))
	ba.EqualJSON(map[string]string{"key": "value"}, "")
}

func TestBytesAssertion_NotEqualJSON(t *testing.T) {
	ba := aa.AssertBytes(t, []byte(`{"key":"value"}`))
	ba.NotEqualJSON(map[string]string{"key": "different"}, "  ")
}
