package aa

import (
	"fmt"
	"strings"
	"testing"
)

type MapAssertion[K comparable, V comparable] struct {
	t *testing.T
	m map[K]V
}

func AssertMap[K comparable, V comparable](t *testing.T, m map[K]V) *MapAssertion[K, V] {
	return &MapAssertion[K, V]{t: t, m: m}
}

func (ma *MapAssertion[K, V]) Equal(b map[K]V) *MapAssertion[K, V] {
	if len(ma.m) != len(b) {
		ma.t.Fatalf("maps are not equal: len(a)=%d, len(b)=%d", len(ma.m), len(b))
	}
	notEqualKeys := make([]K, 0)
	notEqualValues := make([]V, 0)
	for k, v := range ma.m {
		if v != b[k] {
			notEqualKeys = append(notEqualKeys, k)
			notEqualValues = append(notEqualValues, v)
		}
	}
	if len(notEqualKeys) > 0 {
		builder := strings.Builder{}
		for idx, k := range notEqualKeys {
			builder.WriteString("a[")
			builder.WriteString(fmt.Sprint(k))
			builder.WriteString("]=")
			builder.WriteString(fmt.Sprint(notEqualValues[idx]))
			if idx < len(notEqualKeys)-1 {
				builder.WriteString(", ")
			}
		}

		ma.t.Fatalf("maps are not equal: %s", builder.String())
	}

	return ma
}

func (ma *MapAssertion[K, V]) NotEqual(b map[K]V) *MapAssertion[K, V] {
	if len(ma.m) != len(b) {
		return ma
	}
	all := true
	for k, v := range ma.m {
		if v != b[k] {
			all = false
			break
		}
	}
	if all {
		ma.t.Fatalf("maps are equal: a=%v, b=%v", ma.m, b)
	}

	return ma
}

func (ma *MapAssertion[K, V]) ContainsAllKey(key ...K) *MapAssertion[K, V] {
	notContains := make([]K, 0)
	for _, k := range key {
		if _, ok := ma.m[k]; !ok {
			notContains = append(notContains, k)
		}
	}
	if len(notContains) > 0 {
		ma.t.Fatalf("map does not contain key: %v", notContains)
	}

	return ma
}

func (ma *MapAssertion[K, V]) ContainsAllValue(value ...V) *MapAssertion[K, V] {
	candidates := make(map[V]struct{})
	for _, v := range value {
		candidates[v] = struct{}{}
	}
	for _, v := range ma.m {
		if _, ok := candidates[v]; ok {
			delete(candidates, v)
		}
	}
	if len(candidates) > 0 {
		notContains := make([]V, 0)
		for k := range candidates {
			notContains = append(notContains, k)
		}
		ma.t.Fatalf("map does not contain value: %v", notContains)
	}

	return ma
}

func (ma *MapAssertion[K, V]) NotContainsAllKey(key ...K) *MapAssertion[K, V] {
	candidates := make(map[K]struct{})
	for _, k := range key {
		candidates[k] = struct{}{}
	}
	for k := range candidates {
		if _, ok := ma.m[k]; !ok {
			delete(candidates, k)
		}
	}
	if len(candidates) > 0 {
		notContains := make([]K, 0)
		for k := range candidates {
			notContains = append(notContains, k)
		}
		ma.t.Fatalf("map contains key: %v", notContains)
	}

	return ma
}

func (ma *MapAssertion[K, V]) NotContainsAllValue(value ...V) *MapAssertion[K, V] {
	candidates := make(map[V]struct{})
	for _, v := range value {
		candidates[v] = struct{}{}
	}
	for _, v := range ma.m {
		if _, ok := candidates[v]; ok {
			delete(candidates, v)
		}
	}
	if len(candidates) == 0 {
		ma.t.Fatalf("map contains value: %v", value)
	}

	return ma
}

func (ma *MapAssertion[K, V]) ContainsAnyKey(key ...K) *MapAssertion[K, V] {
	for _, k := range key {
		if _, ok := ma.m[k]; ok {
			return ma
		}
	}
	ma.t.Fatalf("map does not contain any key: %v", key)

	return ma
}

func (ma *MapAssertion[K, V]) ContainsAnyValue(value ...V) *MapAssertion[K, V] {
	candidates := make(map[V]struct{})
	for _, v := range value {
		candidates[v] = struct{}{}
	}
	for _, v := range ma.m {
		if _, ok := candidates[v]; ok {
			return ma
		}
	}
	ma.t.Fatalf("map does not contain any value: %v", value)

	return ma
}

func (ma *MapAssertion[K, V]) NotContainsAnyKey(key ...K) *MapAssertion[K, V] {
	for _, k := range key {
		if _, ok := ma.m[k]; ok {
			ma.t.Fatalf("map contains key: %v", k)
		}
	}

	return ma
}

func (ma *MapAssertion[K, V]) NotContainsAnyValue(value ...V) *MapAssertion[K, V] {
	candidates := make(map[V]struct{})
	for _, v := range value {
		candidates[v] = struct{}{}
	}
	for _, v := range ma.m {
		if _, ok := candidates[v]; ok {
			ma.t.Fatalf("map contains value: %v", v)
		}
	}

	return ma
}

func (ma *MapAssertion[K, V]) UniqueValue() *MapAssertion[K, V] {
	unique := make(map[V]struct{})
	dup := make(map[V]struct{})
	for _, v := range ma.m {
		if _, ok := unique[v]; ok {
			dup[v] = struct{}{}
		}
		unique[v] = struct{}{}
	}
	if len(dup) > 0 {
		duplicates := make([]V, 0, len(dup))
		for k := range dup {
			duplicates = append(duplicates, k)
		}
		ma.t.Fatalf("map contains duplicate value: %v", duplicates)
	}

	return ma
}

func (ma *MapAssertion[K, V]) Empty() *MapAssertion[K, V] {
	if len(ma.m) > 0 {
		ma.t.Fatalf("map is not empty: %v", ma.m)
	}

	return ma
}

func (ma *MapAssertion[K, V]) NotEmpty() *MapAssertion[K, V] {
	if len(ma.m) == 0 {
		ma.t.Fatalf("map is empty")
	}

	return ma
}

func (ma *MapAssertion[K, V]) Nil() *MapAssertion[K, V] {
	if ma.m != nil {
		ma.t.Fatalf("map is not nil: %v", ma.m)
	}

	return ma
}

func (ma *MapAssertion[K, V]) NotNil() *MapAssertion[K, V] {
	if ma.m == nil {
		ma.t.Fatalf("map is nil")
	}

	return ma
}

func (ma *MapAssertion[K, V]) LenEqual(length int) *MapAssertion[K, V] {
	if len(ma.m) != length {
		ma.t.Fatalf("map length is not %d: %v", length, ma.m)
	}

	return ma
}

func (ma *MapAssertion[K, V]) LenGreaterThan(length int) *MapAssertion[K, V] {
	if len(ma.m) <= length {
		ma.t.Fatalf("map length is not greater than %d: %v", length, ma.m)
	}

	return ma
}

func (ma *MapAssertion[K, V]) LenNotEqual(length int) *MapAssertion[K, V] {
	if len(ma.m) == length {
		ma.t.Fatalf("map length is %d: %v", length, ma.m)
	}

	return ma
}

func (ma *MapAssertion[K, V]) LenLessThan(length int) *MapAssertion[K, V] {
	if len(ma.m) >= length {
		ma.t.Fatalf("map length is not less than %d: %v", length, ma.m)
	}

	return ma
}
