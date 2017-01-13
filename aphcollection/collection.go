//Package aphcollection contains collection functions for string
//slices
package aphcollection

// Index returns the index of the first instance of s in slice a, or -1 if s is
// not present in a
func Index(a []string, s string) int {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}

// Contains reports whether s is present in slice a
func Contains(a []string, s string) bool {
	return Index(a, s) >= 0
}

// Map applies the given function to each element of a, returning slice of
// results
func Map(a []string, fn func(string) string) []string {
	sl := make([]string, len(a))
	for i, v := range a {
		sl[i] = fn(v)
	}
	return sl
}
