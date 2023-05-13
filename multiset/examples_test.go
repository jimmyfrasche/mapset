package multiset_test

import (
	"fmt"

	"github.com/jimmyfrasche/mapset/multiset"
)

func ExampleOf_Union() {
	x := multiset.Of[string]{"a": 1, "b": 0, "c": 5, "e": 1}
	y := multiset.Of[string]{"c": 3, "d": 4, "e": 9}

	fmt.Println(x.Union(y))

	// Output:
	// map[a:1 c:5 d:4 e:9]
}

func ExampleOf_Intersect() {
	x := multiset.Of[string]{"a": 1, "b": 0, "c": 5, "e": 1}
	y := multiset.Of[string]{"c": 3, "d": 4, "e": 9}

	fmt.Println(x.Intersect(y))

	// Output:
	// map[a:1 c:3 d:4 e:1]
}

func ExampleOf_Add() {
	x := multiset.Of[string]{"a": 1, "b": 0, "c": 5}
	y := multiset.Of[string]{"c": 3, "d": 4}

	fmt.Println(x.Add(y))

	// Output:
	// map[a:1 c:8 d:4]
}
func ExampleOf_Sub() {
	x := multiset.Of[string]{"a": 1, "b": 0, "c": 5}
	y := multiset.Of[string]{"c": 3, "d": 4}

	fmt.Println(x.Sub(y))
	fmt.Println(y.Sub(x))
	// Output:
	// map[a:1 c:2 d:4]
	// map[a:1 d:4]
}

func ExampleOf_Included() {
	x := multiset.Of[string]{"a": 1, "b": 2, "c": 3}
	y := multiset.Of[string]{"a": 1, "b": 3, "c": 4}
	z := multiset.Of[string]{"d": 0, "e": 9, "f": 2}

	fmt.Println("x ≤ y?", x.Included(y))
	fmt.Println("x ≥ y?", y.Included(x))
	fmt.Println("x ≤ z?", x.Included(z))

	// Output:
	// x ≤ y? true
	// x ≥ y? false
	// x ≤ z? false
}

func ExampleOf_ProperIncluded() {
	x := multiset.Of[string]{"a": 1, "b": 2, "c": 3}
	y := multiset.Of[string]{"a": 5, "b": 3, "c": 4}
	z := multiset.Of[string]{"d": 0, "e": 9, "f": 2}

	fmt.Println("x < y?", x.ProperIncluded(y))
	fmt.Println("x > y?", y.ProperIncluded(x))
	fmt.Println("x < z?", x.ProperIncluded(z))

	// Output:
	// x < y? true
	// x > y? false
	// x < z? false
}

func ExampleOf_Equal() {
	x := multiset.Of[string]{"a": 1, "b": 2, "c": 3}
	y := multiset.Of[string]{"a": 1, "b": 3, "c": 4}
	z := multiset.Of[string]{"a": 1, "b": 2, "c": 3, "d": 1}

	fmt.Println("x = y?", x.Equal(y))
	fmt.Println("y = x?", y.Equal(x))
	fmt.Println("x = z?", x.Equal(z))
	fmt.Println("x = x?", x.Equal(x))

	// Output:
	// x = y? false
	// y = x? false
	// x = z? false
	// x = x? true
}

func ExampleOf_Cardinality() {
	x := multiset.Of[string]{"a": 1, "b": 0, "c": 5}
	fmt.Println(x.Cardinality())

	// Output:
	// 6
}
