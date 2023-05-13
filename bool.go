package mapset

func containsBool(b bool) bool {
	return b
}

func mergeBool(lhs, rhs bool) bool {
	return lhs || rhs
}

// Bool provides all relevant set methods with identity as the [ContainsFunc] and disjunction as the [MergeFunc].
type Bool[K comparable] map[K]bool

// Extend adds keys to b.
func (b Bool[K]) Extend(keys ...K) {
	for _, k := range keys {
		b[k] = true
	}
}

// Add key and reports whether it is new.
func (b Bool[K]) Add(key K) bool {
	seen := b[key]
	b[key] = true
	return !seen
}

// Delete removes key and reports if it was present.
func (b Bool[K]) Delete(key K) bool {
	seen := b[key]
	delete(b, key)
	return seen
}

// Remove keys from b.
func (b Bool[K]) Remove(keys ...K) {
	for _, k := range keys {
		delete(b, k)
	}
}

func (b Bool[K]) Union(o Bool[K]) Bool[K] {
	return Union(b, o, containsBool, mergeBool)
}

func (b Bool[K]) Intersect(o Bool[K]) Bool[K] {
	return Intersect(b, o, containsBool, mergeBool)
}

func (b Bool[K]) Diff(o Bool[K]) Bool[K] {
	return Diff(b, o, containsBool)
}

func (b Bool[K]) SymDiff(o Bool[K]) Bool[K] {
	return SymDiff(b, o, containsBool)
}

func (b Bool[K]) Contains(k K) bool {
	return b[k]
}

func (b Bool[K]) Equal(o Bool[K]) bool {
	return Equal(b, o, containsBool)
}

func (b Bool[K]) Subset(o Bool[K]) bool {
	return Subset(b, o, containsBool)
}

func (b Bool[K]) Disjoint(o Bool[K]) bool {
	return Disjoint(b, o, containsBool)
}

func (b Bool[K]) ProperSubset(o Bool[K]) bool {
	return Subset(b, o, containsBool)
}

func (b Bool[K]) Len() int {
	return Len(b, containsBool)
}

func (b Bool[K]) Clone() Bool[K] {
	return Clone(b, containsBool)
}

func (b Bool[K]) Keys() []K {
	return Keys(b, containsBool)
}

func (b Bool[K]) Purge() {
	Purge(b, containsBool)
}
