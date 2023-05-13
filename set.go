package mapset

// Set provides all relevant set methods with nil [ContainsFunc] and [MergeFunc].
type Set[K comparable] map[K]struct{}

// Extend adds keys to s.
func (s Set[K]) Extend(keys ...K) {
	for _, k := range keys {
		s[k] = struct{}{}
	}
}

// Add key and reports whether it is new.
func (s Set[K]) Add(key K) bool {
	seen := s.Contains(key)
	s[key] = struct{}{}
	return !seen
}

// Delete removes key and reports if it was present.
func (s Set[K]) Delete(key K) bool {
	seen := s.Contains(key)
	delete(s, key)
	return seen
}

// Remove keys from s.
func (s Set[K]) Remove(keys ...K) {
	for _, k := range keys {
		delete(s, k)
	}
}

func (s Set[K]) Union(o Set[K]) Set[K] {
	return Union(s, o, nil, nil)
}

func (s Set[K]) Intersect(o Set[K]) Set[K] {
	return Intersect(s, o, nil, nil)
}

func (s Set[K]) Diff(o Set[K]) Set[K] {
	return Diff(s, o, nil)
}

func (s Set[K]) SymDiff(o Set[K]) Set[K] {
	return SymDiff(s, o, nil)
}

func (s Set[K]) Contains(k K) bool {
	_, ok := s[k]
	return ok
}

func (s Set[K]) Equal(o Set[K]) bool {
	return Equal(s, o, nil)
}

func (s Set[K]) Subset(o Set[K]) bool {
	return Subset(s, o, nil)
}

func (s Set[K]) Disjoint(o Set[K]) bool {
	return Disjoint(s, o, nil)
}

func (s Set[K]) ProperSubset(o Set[K]) bool {
	return ProperSubset(s, o, nil)
}
