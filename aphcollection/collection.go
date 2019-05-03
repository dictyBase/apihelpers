//Package aphcollection contains collection functions for string
//slices
package aphcollection

// Remove removes items from the given(a) slice
func Remove(a []string, items ...string) []string {
	var s []string
	for _, v := range items {
		if !Contains(a, v) {
			s = append(s, v)
		}
	}
	return s
}

// Index returns the index of the first instance of s in slice a, or -1 if s is
// not present in a
func Index(a []string, s string) int {
	if len(a) == 0 {
		return -1
	}
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}

// Contains reports whether s is present in slice a
func Contains(a []string, s string) bool {
	if len(a) == 0 {
		return false
	}
	return Index(a, s) >= 0
}

// MapIdx applies the given function to each element of a, returning slice of
// results
func MapIdx(a []string, fn func(string, int) string) []string {
	if len(a) == 0 {
		return a
	}
	sl := make([]string, len(a))
	for i, v := range a {
		sl[i] = fn(v, i)
	}
	return sl
}

// Map applies the given function to each element of a, returning slice of
// results
func Map(a []string, fn func(string) string) []string {
	if len(a) == 0 {
		return a
	}
	sl := make([]string, len(a))
	for i, v := range a {
		sl[i] = fn(v)
	}
	return sl
}
