// Package mapset implements set-like operations for maps.
// This generalizes the existing pattern of map[K]bool using a [ContainsFunc] predicate in place of a bool.
//
// These operations treat maps as sets defined by their keys that happen to have some associated values that are taken along for the ride.
// When multiple values need to be considered for the same key a [MergeFunc] defines what value that key takes.
//
// The [Set] and [Bool] types package these operations into types that can be used by conversion from a map[K]struct{} or map[K]bool, respectively.
package mapset
