package aa

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type SliceAssertion[T comparable] struct {
	t *testing.T
	a []T
}

func AssertSlice[T comparable](t *testing.T, a []T) *SliceAssertion[T] {
	return &SliceAssertion[T]{t: t, a: a}
}

func (sa *SliceAssertion[T]) Equal(b []T) *SliceAssertion[T] {
	if len(sa.a) != len(b) {
		sa.t.Fatalf("slices are not equal: len(a)=%d, len(b)=%d", len(sa.a), len(b))
	}
	notEqualIndexed := make([]int, 0)
	notEqualValues := make([]T, 0)
	for i, v := range sa.a {
		if v != b[i] {
			notEqualIndexed = append(notEqualIndexed, i)
			notEqualValues = append(notEqualValues, v)
		}
	}
	if len(notEqualIndexed) > 0 {
		builder := strings.Builder{}
		for idx, i := range notEqualIndexed {
			builder.WriteString("a[")
			builder.WriteString(strconv.Itoa(i))
			builder.WriteString("]=")
			builder.WriteString(fmt.Sprint(notEqualValues[idx]))
			if idx < len(notEqualIndexed)-1 {
				builder.WriteString(", ")
			}
		}

		sa.t.Fatalf("slices are not equal: %s", builder.String())
	}

	return sa
}

func (sa *SliceAssertion[T]) NotEqual(b []T) *SliceAssertion[T] {
	if len(sa.a) != len(b) {
		return sa
	}
	all := true
	for i, v := range sa.a {
		if v != b[i] {
			all = false
			break
		}
	}
	if all {
		sa.t.Fatalf("slices are equal: a=%v, b=%v", sa.a, b)
	}

	return sa
}

func (sa *SliceAssertion[T]) Contains(v T) *SliceAssertion[T] {
	contains := false
	for _, e := range sa.a {
		if e == v {
			contains = true
			break
		}
	}
	if !contains {
		sa.t.Fatalf("slice does not contain value: %v", v)
	}

	return sa
}

func (sa *SliceAssertion[T]) NotContains(v T) *SliceAssertion[T] {
	contains := false
	for _, e := range sa.a {
		if e == v {
			contains = true
			break
		}
	}
	if contains {
		sa.t.Fatalf("slice contains value: %v", v)
	}

	return sa
}

func (sa *SliceAssertion[T]) ContainsAll(v ...T) *SliceAssertion[T] {
	candidates := make(map[T]struct{})
	for _, e := range v {
		candidates[e] = struct{}{}
	}
	for _, e := range sa.a {
		if _, ok := candidates[e]; ok {
			delete(candidates, e)
		}
	}
	if len(candidates) > 0 {
		values := make([]T, 0, len(candidates))
		for e := range candidates {
			values = append(values, e)
		}
		sa.t.Fatalf("slice does not contain values: %v", values)
	}

	return sa
}

func (sa *SliceAssertion[T]) NotContainsAll(v ...T) *SliceAssertion[T] {
	candidates := make(map[T]struct{})
	for _, e := range v {
		candidates[e] = struct{}{}
	}
	for _, e := range sa.a {
		if _, ok := candidates[e]; ok {
			delete(candidates, e)
		}
	}
	if len(candidates) == 0 {
		sa.t.Fatalf("slice contains all values: %v", v)
	}

	return sa
}

func (sa *SliceAssertion[T]) ContainsAny(v ...T) *SliceAssertion[T] {
	candidates := make(map[T]struct{})
	for _, e := range v {
		candidates[e] = struct{}{}
	}
	for _, e := range sa.a {
		if _, ok := candidates[e]; ok {
			return sa
		}
	}
	values := make([]T, 0, len(candidates))
	for e := range candidates {
		values = append(values, e)
	}
	sa.t.Fatalf("slice does not contain any values: %v", values)

	return sa
}

func (sa *SliceAssertion[T]) NotContainsAny(v ...T) *SliceAssertion[T] {
	candidates := make(map[T]struct{})
	for _, e := range v {
		candidates[e] = struct{}{}
	}
	contains := make([]T, 0, len(candidates))
	for _, e := range sa.a {
		if _, ok := candidates[e]; ok {
			contains = append(contains, e)
		}
	}
	if len(contains) > 0 {
		sa.t.Fatalf("slice contains values: %v", contains)
	}

	return sa
}

func (sa *SliceAssertion[T]) Unique() *SliceAssertion[T] {
	m := make(map[T]struct{})
	d := make(map[T]struct{})
	for _, e := range sa.a {
		if _, ok := m[e]; ok {
			d[e] = struct{}{}
		}
		m[e] = struct{}{}
	}
	if len(d) > 0 {
		values := make([]T, 0, len(d))
		for e := range d {
			values = append(values, e)
		}
		sa.t.Fatalf("slice contains duplicate values: %v", values)
	}

	return sa
}

func (sa *SliceAssertion[T]) Same(b []T) *SliceAssertion[T] {
	if len(sa.a) != len(b) {
		sa.t.Fatalf("slices are not same: len(a)=%d, len(b)=%d", len(sa.a), len(b))
		return sa
	}
	if cap(sa.a) != cap(b) {
		sa.t.Fatalf("slices are not same: cap(a)=%d, cap(b)=%d", cap(sa.a), cap(b))
		return sa
	}
	notSameIndexes := make([]int, 0, len(sa.a))
	notSameValues := make([]T, 0, len(sa.a))
	for i, v := range sa.a {
		if v != b[i] {
			notSameIndexes = append(notSameIndexes, i)
			notSameValues = append(notSameValues, v)
		}
	}
	if len(notSameIndexes) > 0 {
		builder := strings.Builder{}
		for idx, i := range notSameIndexes {
			builder.WriteString("a[")
			builder.WriteString(strconv.Itoa(i))
			builder.WriteString("]=")
			builder.WriteString(fmt.Sprint(notSameValues[i]))
			if idx < len(notSameIndexes)-1 {
				builder.WriteString(", ")
			}
		}

		sa.t.Fatalf("slices are not same: %s", builder.String())
	}

	return sa
}

func (sa *SliceAssertion[T]) Empty() *SliceAssertion[T] {
	if len(sa.a) > 0 {
		sa.t.Fatalf("slice is not empty: %v", sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) Nil() *SliceAssertion[T] {
	if sa.a != nil {
		sa.t.Fatalf("slice is not nil: %v", sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) NotEmpty() *SliceAssertion[T] {
	if len(sa.a) == 0 {
		sa.t.Fatalf("slice is empty")
	}

	return sa
}

func (sa *SliceAssertion[T]) NotNil() *SliceAssertion[T] {
	if sa.a == nil {
		sa.t.Fatalf("slice is nil")
	}

	return sa
}

func (sa *SliceAssertion[T]) LenEqual(l int) *SliceAssertion[T] {
	if len(sa.a) != l {
		sa.t.Fatalf("slice length is not %d: %v", l, sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) LenNotEqual(l int) *SliceAssertion[T] {
	if len(sa.a) == l {
		sa.t.Fatalf("slice length is %d: %v", l, sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) LenGreaterThan(l int) *SliceAssertion[T] {
	if len(sa.a) <= l {
		sa.t.Fatalf("slice length is not greater than %d: %v", l, sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) LenLessThan(l int) *SliceAssertion[T] {
	if len(sa.a) >= l {
		sa.t.Fatalf("slice length is not less than %d: %v", l, sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) CapEqual(c int) *SliceAssertion[T] {
	if cap(sa.a) != c {
		sa.t.Fatalf("slice capacity is not %d: %v", c, sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) CapNotEqual(c int) *SliceAssertion[T] {
	if cap(sa.a) == c {
		sa.t.Fatalf("slice capacity is %d: %v", c, sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) CapGreaterThan(c int) *SliceAssertion[T] {
	if cap(sa.a) <= c {
		sa.t.Fatalf("slice capacity is not greater than %d: %v", c, sa.a)
	}

	return sa
}

func (sa *SliceAssertion[T]) CapLessThan(c int) *SliceAssertion[T] {
	if cap(sa.a) >= c {
		sa.t.Fatalf("slice capacity is not less than %d: %v", c, sa.a)
	}

	return sa
}
