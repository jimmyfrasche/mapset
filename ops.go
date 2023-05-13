package mapset

// Union creates a new M that is the union of lhs and rhs.
// For an item of either map to be included the value of the item must pass the [ContainsFunc] check.
// When an item exists in both maps the pair of values is merged and added to the result map
// if the merged value passes the [ContainsFunc] check.
func Union[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V], merge MergeFunc[V]) M {
	out := M{}
	for k, v := range lhs {
		if contains.Check(v) {
			out[k] = v
		}
	}
	for k, v := range rhs {
		if contains.Check(v) {
			if L, ok := out[k]; !ok {
				// new key, add
				out[k] = v
			} else {
				// key exists in both maps, do a merge
				v := merge.Into(L, v)

				out[k] = v
				// make sure two valid entries aren't merged into an invalid entry
				if !contains.Check(v) {
					delete(out, k)
				}
			}
		}
	}
	return out
}

// Intersect creates a new M that is the intersection of lhs and rhs.
// All pairs values pass [ContainsFunc], are merged, and the result passes the [ContainsFunc] for inclusion.
func Intersect[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V], merge MergeFunc[V]) M {
	out := M{}
	for k, v := range lhs {
		vp, ok := rhs[k]
		if ok && contains.Check(v) && contains.Check(vp) {
			vf := merge.Into(v, vp)
			if contains.Check(vf) {
				out[k] = vf
			}
		}
	}
	return out
}

// Diff is the set difference of lhs and rhs.
// The result contains all items in lhs that pass the [ContainsFunc] check that are not in rhs.
// Diff is also known as: relative complement, kick out, or except.
func Diff[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V]) M {
	out := M{}
	for k, v := range lhs {
		if contains.Check(v) && !Contains(rhs, k, contains) {
			out[k] = v
		}
	}
	return out
}

// SymDiff is the symmetric difference of lhs and rhs.
// The results contains the values in lhs and rhs but not in both, according to the [ContainsFunc] check.
func SymDiff[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V]) M {
	out := M{}
	for k, v := range lhs {
		// k in lhs but not rhs
		if contains.Check(v) && !Contains(rhs, k, contains) {
			out[k] = v
		}
	}
	for k, v := range rhs {
		// k in rhs but not lhs
		if contains.Check(v) && !Contains(lhs, k, contains) {
			out[k] = v
		}
	}
	return out
}

// Disjoint returns true if no keys of lhs are present in rhs.
func Disjoint[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V]) bool {
	for k, v := range lhs {
		if contains.Check(v) && Contains(rhs, k, contains) {
			return false
		}
	}
	return true
}

func subset[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V]) (int, bool) {
	len := 0
	for k, v := range lhs {
		if contains.Check(v) && !Contains(rhs, k, contains) {
			return 0, false
		}
		len++
	}
	return len, true
}

// Equal tests equality of the keys whose values pass the [ContainsFunc] check. The values are otherwise not considered.
func Equal[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V]) bool {
	m, ok := subset(lhs, rhs, contains)
	if !ok {
		return false
	}
	n, ok := subset(rhs, lhs, contains)
	return ok && m == n
}

// Subset tests whether lhs is a Subset of rhs.
func Subset[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V]) bool {
	_, ok := subset(lhs, rhs, contains)
	return ok
}

// ProperSubset checks that lhs Subset rhs but not lhs Equal rhs.
func ProperSubset[K comparable, V any, M ~map[K]V](lhs, rhs M, contains ContainsFunc[V]) bool {
	n, ok := subset(lhs, rhs, contains)
	if !ok {
		return false
	}
	// can only be equal
	if len(rhs) == n {
		return false
	}
	_, ok = subset(rhs, lhs, contains)
	return !ok
}

// Len counts the keys in m whose values pass the [ContainsFunc] check.
func Len[K comparable, V any, M ~map[K]V](m M, contains ContainsFunc[V]) int {
	var n int
	for _, v := range m {
		if contains.Check(v) {
			n++
		}
	}
	return n
}

// Clone returns a copy of M without items that fail the [ContainsFunc] check.
func Clone[K comparable, V any, M ~map[K]V](m M, contains ContainsFunc[V]) M {
	out := M{}
	for k, v := range m {
		if contains.Check(v) {
			out[k] = v
		}
	}
	return out
}

// Keys collects all the the keys whose corresponding values pass the [ContainsFunc] check.
func Keys[K comparable, V any, M ~map[K]V](m M, contains ContainsFunc[V]) []K {
	var out []K
	for k, v := range m {
		if contains.Check(v) {
			out = append(out, k)
		}
	}
	return out
}

// Values collects all the values in m that pass the [ContainsFunc] check.
func Values[K comparable, V any, M ~map[K]V](m M, contains ContainsFunc[V]) []V {
	var out []V
	for _, v := range m {
		if contains.Check(v) {
			out = append(out, v)
		}
	}
	return out
}

// Purge removes all keys who fail the [ContainsFunc] check.
func Purge[K comparable, V any, M ~map[K]V](m M, contains ContainsFunc[V]) {
	if contains == nil {
		return
	}
	for k, v := range m {
		if !contains.Check(v) {
			delete(m, k)
		}
	}
}
