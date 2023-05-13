package mapset

// MergeFunc takes two V and merges them into a single value.
// A nil MergeFunc always returns the lhs value.
//
// MergeFunc is used for combining the values of a key that exists in two maps that are being combined in some way.
//
// When used with a [ContainsFunc], both lhs and rhs will have passed the check before the MergeFunc is called and the resulting value is also checked.
type MergeFunc[V any] func(lhs, rhs V) V

// Into returns m(lhs, rhs) or lhs if m == nil.
func (m MergeFunc[V]) Into(lhs, rhs V) V {
	if m == nil {
		return lhs
	}
	return m(lhs, rhs)
}
