package mapset_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/jimmyfrasche/mapset"
)

func ExampleUnion() {
	x := map[string]int{"a": 0, "b": 1, "d": -5}
	y := map[string]int{"b": 2, "c": 1, "d": 5}

	contains := func(n int) bool {
		return n != 0
	}

	merge := func(a, b int) int {
		return a + b
	}

	// our contains func discards "a" from x
	// our merge func makes "b" 3 in the resulting map
	// For "d" the merge func results in a value that does not pass the contains check,
	// so it is not included in z.

	z := mapset.Union(x, y, contains, merge)
	fmt.Println("z:", z)

	// with the default contains and merge, no items are discarded and the values
	// for keys in both default to the values in x.
	zp := mapset.Union(x, y, nil, nil)
	fmt.Println("zp:", zp)

	// Output:
	// z: map[b:3 c:1]
	// zp: map[a:0 b:1 c:1 d:-5]
}

func ExampleIntersect() {
	x := map[string]int{"a": 0, "b": 1, "d": -5}
	y := map[string]int{"b": 2, "c": 1, "d": 5}

	z := mapset.Intersect(x, y, nil, nil)
	fmt.Println(z)

	// Output:
	// map[b:1 d:-5]
}

func ExampleDiff() {
	x := map[string]int{"a": 0, "b": 1, "d": -5}
	y := map[string]int{"b": 2, "c": 1, "d": 5}

	z := mapset.Diff(x, y, nil)
	fmt.Println(z)

	// Output:
	// map[a:0]
}

func ExampleSymDiff() {
	x := map[string]int{"a": 0, "b": 1, "d": -5}
	y := map[string]int{"b": 2, "c": 1, "d": 5, "e": 7}

	contains := func(n int) bool {
		return n >= 0
	}

	// z has d:5 because the "d" in x is skipped by contains.Check.
	z := mapset.SymDiff(x, y, contains)
	fmt.Println("z:", z)

	zp := mapset.SymDiff(x, y, nil)
	fmt.Println("zp:", zp)

	// Output:
	// z: map[a:0 c:1 d:5 e:7]
	// zp: map[a:0 c:1 e:7]
}

func ExampleDisjoint() {
	x := map[string]int{"a": 0, "b": 1, "d": -5}
	y := map[string]int{"b": 1, "e": 5, "f": 7}

	fmt.Println("nil contains:", mapset.Disjoint(x, y, nil))

	// b: 1 is the only kv-pair in both maps
	// so this leaves them with disjoint key-sets
	contains := func(n int) bool {
		return n != 1
	}

	fmt.Println("contains removes 1s:", mapset.Disjoint(x, y, contains))

	// Output:
	// nil contains: false
	// contains removes 1s: true
}

func ExampleEqual() {
	a := map[string]int{"a": 4, "b": 7, "c": -1}
	b := map[string]int{"a": 3, "b": 2, "c": 11}
	c := map[string]int{"b": 2, "c": 11, "d": 86}

	if mapset.Equal(a, b, nil) {
		fmt.Println("a = b")
	}

	if !mapset.Equal(a, c, nil) {
		fmt.Println("a != c")
	}

	contains := func(x int) bool {
		return x >= 0
	}

	if !mapset.Equal(a, b, contains) {
		fmt.Println("a != b as contains removes a[c]")
	}

	// Output:
	// a = b
	// a != c
	// a != b as contains removes a[c]
}

func ExampleSubset() {
	a := map[string]int{"a": 4, "b": 7}
	b := map[string]int{"a": 3, "b": 2, "c": 1}

	if mapset.Subset(a, b, nil) {
		fmt.Println("the keys of a are a subset of the keys of b")
	}

	if mapset.Subset(a, a, nil) {
		fmt.Println("a is a subset of a")
	}

	if !mapset.Subset(b, a, nil) {
		fmt.Println("the keys of b are not a subset of the keys of a")
	}

	// Output:
	// the keys of a are a subset of the keys of b
	// a is a subset of a
	// the keys of b are not a subset of the keys of a
}

func ExampleProperSubset() {
	a := map[string]int{"a": 4, "b": 7}
	b := map[string]int{"a": 3, "b": 2, "c": 1}

	if mapset.ProperSubset(a, b, nil) {
		fmt.Println("the keys of a are a proper subset of the keys of b")
	}

	if !mapset.ProperSubset(a, a, nil) {
		fmt.Println("a is not a proper subset of a")
	}

	if !mapset.ProperSubset(b, a, nil) {
		fmt.Println("the keys of b are not a proper subset of the keys of a")
	}

	// Output:
	// the keys of a are a proper subset of the keys of b
	// a is not a proper subset of a
	// the keys of b are not a proper subset of the keys of a
}

func TestLen(t *testing.T) {
	x := map[string]int{"a": 0, "b": 1, "d": -5}

	if len(x) != mapset.Len(x, nil) {
		t.Fatal("nil contains check must be same as len(m)")
	}

	n := mapset.Len(x, func(int) bool { return false })
	if n != 0 {
		t.Fatal("no items should pass contains func")
	}
}

func ExampleClone() {
	x := map[string]int{"a": -1, "b": 0, "c": 1}

	contains := func(n int) bool {
		return n >= 0
	}

	// this is the same behavior of maps.Clone
	y := mapset.Clone(x, nil)

	// this clone will not contain "a" as -1 fails the contains check
	z := mapset.Clone(x, contains)

	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Output:
	// y: map[a:-1 b:0 c:1]
	// z: map[b:0 c:1]
}

func ExampleKeys() {
	m := map[string]bool{"a": true, "b": true, "c": false, "d": true}

	contains := func(b bool) bool {
		return b
	}

	k := mapset.Keys(m, contains)

	sort.Strings(k)

	fmt.Println("k:", k)

	// Output:
	// k: [a b d]
}

func ExampleValues() {
	m := map[int]string{0: "a", 1: "b", 2: "c"}

	contains := func(s string) bool {
		return s != "a"
	}

	v := mapset.Values(m, contains)

	sort.Strings(v)

	fmt.Println("v:", v)
	//Output:
	// v: [b c]
}

func ExamplePurge() {
	x := map[string]int{"a": -1, "b": 0, "c": 1}

	// this is a noop as there is nothing to purge
	mapset.Purge(x, nil)

	fmt.Println("nil contains:", x)

	contains := func(n int) bool {
		return n >= 0
	}

	// this deleted "a" from x
	mapset.Purge(x, contains)

	fmt.Println("nonnil contains:", x)

	// Output:
	// nil contains: map[a:-1 b:0 c:1]
	// nonnil contains: map[b:0 c:1]
}
