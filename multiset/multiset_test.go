package multiset

import (
	"math"
	"testing"
)

func TestInc(t *testing.T) {
	m := Of[int]{0: 1, 1: 2}

	m.Inc(0, 1)
	if m[0] != 2 {
		t.Fatal("m[0] + 1 ≠ 2")
	}

	if m.Inc(1, 0) != 2 {
		t.Fatal("m[1] + 0 ≠ m[1]")
	}

	m.Inc(3, 1)
	if !m.Contains(3) {
		t.Fatal("did not insert m[3]")
	}

	defer func() {
		x := recover()
		if x == nil {
			t.Fatal("did not panic")
		}
		if x != ErrOverflow {
			t.Fatal("did not panic correctly")
		}
	}()
	m.Inc(0, math.MaxUint64)
}

func TestDec(t *testing.T) {
	m := Of[int]{}

	if m.Dec(0, 0) != 0 {
		t.Fatal("if key does not exist, should return 0")
	}

	m[0] = 2

	if m.Dec(0, 0) != 2 {
		t.Fatal("decrement of 0 should return m[k]")
	}

	if m.Dec(0, 1) != 1 {
		t.Fatal("2 - 1 ≠ 1")
	}

	m.Dec(0, 1)
	_, ok := m[0]
	if ok {
		t.Fatal("Dec should remove key if multiplicity hits 0")
	}
}
