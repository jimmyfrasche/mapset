package mapset

// ContainsFunc decides if an item belongs to a set based on its value.
// A nil ContainsFunc always returns true.
//
// ContainsFunc generalizes the pattern of letting a map[K]bool contain an item only if m[k] is true.
//
// See [Contains].
type ContainsFunc[V any] func(V) bool

// Check calls c with v or returns true if c is nil.
func (c ContainsFunc[V]) Check(v V) bool {
	if c == nil {
		return true
	}
	return c(v)
}

// Contains returns false if key is not in the set m.
//
// If the key is not in the map m, Contains returns false.
//
// If the key is in the map, Contains returns the result of contains.Check.
//
// Everything in this package that accepts a ContainsFunc applies the same logic.
func Contains[K comparable, V any, M ~map[K]V](m M, key K, contains ContainsFunc[V]) bool {
	v, ok := m[key]
	if !ok {
		return false
	}
	return contains.Check(v)
}
