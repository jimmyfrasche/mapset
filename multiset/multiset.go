// Package multiset uses the primitives in mapset to implement a simple multiset.
package multiset

import (
	"errors"
	"math/bits"

	"github.com/jimmyfrasche/mapset"
)

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint64) uint64 {
	if a < b {
		return b
	}
	return a
}

func properSubtraction(a, b uint64) uint64 {
	if a < b {
		return 0
	}
	return a - b
}

// ErrOverflow is used whenever addition of two uint64 would overflow.
// This is an unlikely thing that few will ever need to consider so it is always reported via
//
//	panic(ErrOverflow)
var ErrOverflow = errors.New("overflow")

func sum(a, b uint64) uint64 {
	v, over := bits.Add64(a, b, 0)
	if over != 0 {
		panic(ErrOverflow)
	}
	return v
}

func in(n uint64) bool {
	return n > 0
}

func lt(x, y uint64) bool {
	return x < y
}

func lte(x, y uint64) bool {
	return x <= y
}

func eq(x, y uint64) bool {
	return x == y
}

// multiset.Of[K] is a set of K where keys can be stored more than once.
// The number of times a key is stored, m[k], is called the multiplicity of the key.
//
// A multiplicity of 0 is the same as not being in the set.
type Of[K comparable] map[K]uint64

// Union chooses the maximum of multiplicities of both multisets.
//
//	r[k] = max(m[k], o[k])
func (m Of[K]) Union(o Of[K]) Of[K] {
	return mapset.Union(m, o, in, max)
}

// Intersect chooses the minimum of multiplicities of both multisets.
//
//	r[k] = min(m[k], o[k])
func (m Of[K]) Intersect(o Of[K]) Of[K] {
	return mapset.Union(m, o, in, min)
}

// Add chooses the sum of multiplicities of both multisets.
// It panics if any sum overflows.
//
//	r[k] = m[k] + o[k]
func (m Of[K]) Add(o Of[K]) Of[K] {
	return mapset.Union(m, o, in, sum)
}

// Sub chooses the proper subtraction of both multisets.
//
//	r[k] = max(m[k] - o[k], 0)
func (m Of[K]) Sub(o Of[K]) Of[K] {
	return mapset.Union(m, o, in, properSubtraction)
}

// Inc adds x to m[k]. This panics if the addition overflows.
func (m Of[K]) Inc(k K, x uint64) uint64 {
	y := m[k]
	if x == 0 {
		return y
	}
	v := sum(y, x)
	m[k] = v
	return v
}

// Dec properly subtracts x from m[k].
func (m Of[K]) Dec(k K, x uint64) uint64 {
	y, ok := m[k]
	if !ok {
		return 0
	}
	if x == 0 {
		return y
	}
	v := properSubtraction(y, x)
	if v == 0 {
		delete(m, k)
		return 0
	}
	m[k] = v
	return v
}

// If n >= 0, call Inc(k, uint(n)); otherwise Dec(k, uint64(-n)).
func (m Of[K]) IncDec(k K, n int) uint64 {
	switch {
	case n > 0:
		return m.Inc(k, uint64(n))
	case n < 0:
		return m.Dec(k, uint64(-n))
	}
	// n == 0 is a noop so just return multiplicity of k.
	return m[k]
}

// Contains k if m[k] > 0.
func (m Of[K]) Contains(k K) bool {
	return m[k] > 0
}

func (m Of[K]) rel(o Of[K], rel func(x, y uint64) bool) bool {
	for k, mv := range m {
		ov := o[k]
		if !rel(mv, ov) {
			return false
		}
	}
	for k, ov := range o {
		mv := m[k]
		if !rel(mv, ov) {
			return false
		}
	}
	return true
}

// m Included in o if the multiplicity of each item in m is <= the corresponding multiplicity in o.
//
// true if, for all k:
//
//	m[k] <= o[k]
func (m Of[K]) Included(o Of[K]) bool {
	return m.rel(o, lte)
}

// m is ProperIncluded in o if the multiplicity of each item in m is < the corresponding multiplicity in o.
//
// true if, for all k:
//
//	m[k] < o[k]
func (m Of[K]) ProperIncluded(o Of[K]) bool {
	return m.rel(o, lt)
}

// true if, for all k:
//
//	m[k] == o[k]
func (m Of[K]) Equal(o Of[K]) bool {
	return m.rel(o, eq)
}

// Cardinality is the sum of the multiplicities of all items.
// It panics if the sum overflows.
func (m Of[K]) Cardinality() uint64 {
	var n uint64
	for _, v := range m {
		n = sum(n, v)
	}
	return n
}

// Keys returns the support of m as a slice.
func (m Of[K]) Keys() []K {
	return mapset.Keys(m, in)
}

func (m Of[K]) Clone() Of[K] {
	return mapset.Clone(m, in)
}

// Purge removes all keys of multiplicity 0.
func (m Of[K]) Purge() {
	mapset.Purge(m, in)
}
