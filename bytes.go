package aa

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

type Bytes interface {
	[]byte | string
}

type BytesAssertion[T Bytes] struct {
	t *testing.T
	a T
}

func AssertBytes[T Bytes](t *testing.T, a T) *BytesAssertion[T] {
	return &BytesAssertion[T]{
		t: t,
		a: a,
	}
}

func (ba *BytesAssertion[T]) Equal(expected T) *BytesAssertion[T] {
	if len(ba.a) != len(expected) {
		ba.t.Fatalf("bytes are not equal: len(a)=%d, len(b)=%d", len(ba.a), len(expected))
	}
	switch v := any(ba.a).(type) {
	case []byte:
		e := any(expected).([]byte)
		for i := range v {
			if v[i] != e[i] {
				ba.t.Fatalf("bytes are not equal: a[%d]=%d, b[%d]=%d", i, v[i], i, e[i])
			}
		}
	case string:
		e := any(expected).(string)
		for i := range v {
			if v[i] != e[i] {
				ba.t.Fatalf("bytes are not equal: a[%d]=%d, b[%d]=%d", i, v[i], i, e[i])
			}
		}
	default:
		ba.t.Fatalf("unexpected type: %v", reflect.TypeOf(ba.a))
	}
	return ba
}

func (ba *BytesAssertion[T]) NotEqual(expected T) *BytesAssertion[T] {
	if len(ba.a) != len(expected) {
		return ba
	}
	all := true
	switch v := any(ba.a).(type) {
	case []byte:
		e := any(expected).([]byte)
		for i := range v {
			if v[i] != e[i] {
				all = false
				break
			}
		}
	case string:
		e := any(expected).(string)
		for i := range v {
			if v[i] != e[i] {
				all = false
				break
			}
		}
	}
	if all {
		ba.t.Fatalf("bytes are equal")
	}
	return ba
}

func (ba *BytesAssertion[T]) IsEqualFold(b []byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		if !bytes.EqualFold(a, b) {
			ba.t.Fatalf("bytes are not equal fold: a=%v, b=%v", a, b)
		}
	case string:
		if !bytes.EqualFold([]byte(a), b) {
			ba.t.Fatalf("bytes are not equal fold: a=%v, b=%v", a, b)
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) IsNotEqualFold(b []byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		if bytes.EqualFold(a, b) {
			ba.t.Fatalf("bytes are equal fold: a=%v, b=%v", a, b)
		}
	case string:
		if bytes.EqualFold([]byte(a), b) {
			ba.t.Fatalf("bytes are equal fold: a=%v, b=%v", a, b)
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) Contains(v []byte) *BytesAssertion[T] {
	contains := false
	switch a := any(ba.a).(type) {
	case []byte:
		if bytes.Contains(a, v) {
			contains = true
		}
	case string:
		if bytes.Contains([]byte(a), v) {
			contains = true
		}
	}
	if !contains {
		ba.t.Fatalf("bytes does not contain value: %v", v)
	}
	return ba
}

func (ba *BytesAssertion[T]) NotContains(v []byte) *BytesAssertion[T] {
	contains := false
	switch a := any(ba.a).(type) {
	case []byte:
		if !bytes.Contains(a, v) {
			contains = true
		}
	case string:
		if !bytes.Contains([]byte(a), v) {
			contains = true
		}
	}
	if !contains {
		ba.t.Fatalf("bytes contains value: %v", v)
	}
	return ba
}

func (ba *BytesAssertion[T]) ContainsAll(v ...[]byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		for _, e := range v {
			if !bytes.Contains(a, e) {
				ba.t.Fatalf("bytes does not contain value: %v", e)
			}
		}
	case string:
		for _, e := range v {
			if !bytes.Contains([]byte(a), e) {
				ba.t.Fatalf("bytes does not contain value: %v", e)
			}
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) NotContainsAll(v ...[]byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		for _, e := range v {
			if bytes.Contains(a, e) {
				ba.t.Fatalf("bytes contains value: %v", e)
			}
		}
	case string:
		for _, e := range v {
			if bytes.Contains([]byte(a), e) {
				ba.t.Fatalf("bytes contains value: %v", e)
			}
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) ContainsAny(v ...[]byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		for _, e := range v {
			if bytes.Contains(a, e) {
				return ba
			}
		}
		ba.t.Fatalf("bytes does not contain any value: %v", v)
	case string:
		for _, e := range v {
			if bytes.Contains([]byte(a), e) {
				return ba
			}
		}
		ba.t.Fatalf("bytes does not contain any value: %v", v)
	}
	return ba
}

func (ba *BytesAssertion[T]) NotContainsAny(v ...[]byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		for _, e := range v {
			if bytes.Contains(a, e) {
				ba.t.Fatalf("bytes contains value: %v", e)
			}
		}
	case string:
		for _, e := range v {
			if bytes.Contains([]byte(a), e) {
				ba.t.Fatalf("bytes contains value: %v", e)
			}
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) IsEmpty() *BytesAssertion[T] {
	if len(ba.a) > 0 {
		ba.t.Fatalf("bytes is not empty: %v", ba.a)
	}
	return ba
}

func (ba *BytesAssertion[T]) IsNotEmpty() *BytesAssertion[T] {
	if len(ba.a) == 0 {
		ba.t.Fatalf("bytes is empty")
	}
	return ba
}

func (ba *BytesAssertion[T]) IsNil() *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		if a != nil {
			ba.t.Fatalf("bytes is not nil: %v", a)
		}
	case string:
		ba.t.Fatalf("bytes is not nil: %v", a)
	}
	return ba
}

func (ba *BytesAssertion[T]) IsNotNil() *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		if a == nil {
			ba.t.Fatalf("bytes is nil")
		}
	case string:
	}
	return ba
}

func (ba *BytesAssertion[T]) HasPrefix(prefix []byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		if !bytes.HasPrefix(a, prefix) {
			ba.t.Fatalf("bytes does not have prefix: %v", prefix)
		}
	case string:
		if !bytes.HasPrefix([]byte(a), prefix) {
			ba.t.Fatalf("bytes does not have prefix: %v", prefix)
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) HasSuffix(suffix []byte) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		if !bytes.HasSuffix(a, suffix) {
			ba.t.Fatalf("bytes does not have suffix: %v", suffix)
		}
	case string:
		if !bytes.HasSuffix([]byte(a), suffix) {
			ba.t.Fatalf("bytes does not have suffix: %v", suffix)
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) LenEqual(l int) *BytesAssertion[T] {
	if len(ba.a) != l {
		ba.t.Fatalf("bytes length is not %d: %v", l, ba.a)
	}
	return ba
}

func (ba *BytesAssertion[T]) LenNotEqual(l int) *BytesAssertion[T] {
	if len(ba.a) == l {
		ba.t.Fatalf("bytes length is %d: %v", l, ba.a)
	}
	return ba
}

func (ba *BytesAssertion[T]) LenGreaterThan(l int) *BytesAssertion[T] {
	if len(ba.a) <= l {
		ba.t.Fatalf("bytes length is not greater than %d: %v", l, ba.a)
	}
	return ba
}

func (ba *BytesAssertion[T]) LenLessThan(l int) *BytesAssertion[T] {
	if len(ba.a) >= l {
		ba.t.Fatalf("bytes length is not less than %d: %v", l, ba.a)
	}
	return ba
}

func (ba *BytesAssertion[T]) EqualHex(expected string) *BytesAssertion[T] {
	switch a := any(ba.a).(type) {
	case []byte:
		if hex.EncodeToString(a) != expected {
			ba.t.Fatalf("bytes are not equal hex: a=%v, b=%v", a, expected)
		}
	case string:
		ba.t.Fatalf("bytes are not equal hex: a=%v, b=%v", a, expected)
	}
	return ba
}

const (
	DefaultJsonIndent = "  "
)

func (ba *BytesAssertion[T]) EqualJSON(value any, indent string) *BytesAssertion[T] {
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", indent)
	_ = encoder.Encode(value)
	data := buffer.Bytes()

	switch a := any(ba.a).(type) {
	case []byte:
		if !bytes.Equal(a, data) {
			ba.t.Fatalf("bytes are not equal json: a=%v, b=%v", a, data)
		}
	case string:
		if !bytes.Equal([]byte(a), data) {
			ba.t.Fatalf("bytes are not equal json: a=%v, b=%v", a, data)
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) NotEqualJSON(value any, indent string) *BytesAssertion[T] {
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", indent)
	_ = encoder.Encode(value)
	data := buffer.Bytes()

	switch a := any(ba.a).(type) {
	case []byte:
		if bytes.Equal(a, data) {
			ba.t.Fatalf("bytes are equal json: a=%v, b=%v", a, data)
		}
	case string:
		if bytes.Equal([]byte(a), data) {
			ba.t.Fatalf("bytes are equal json: a=%v, b=%v", a, data)
		}
	}
	return ba
}

func (ba *BytesAssertion[T]) EqualProtoBuf(value proto.Message) *BytesAssertion[T] {
	data, _ := proto.Marshal(value)

	switch a := any(ba.a).(type) {
	case []byte:
		if !bytes.Equal(a, data) {
			ba.t.Fatalf("bytes are not equal protobuf: a=%v, b=%v", a, data)
		}
	case string:
		ba.t.Fatalf("bytes are not equal protobuf: a=%v, b=%v", a, data)
	}
	return ba
}

func (ba *BytesAssertion[T]) NotEqualProtoBuf(value proto.Message) *BytesAssertion[T] {
	data, _ := proto.Marshal(value)

	switch a := any(ba.a).(type) {
	case []byte:
		if bytes.Equal(a, data) {
			ba.t.Fatalf("bytes are equal protobuf: a=%v, b=%v", a, data)
		}
	case string:
		ba.t.Fatalf("bytes are equal protobuf: a=%v, b=%v", a, data)
	}
	return ba
}
