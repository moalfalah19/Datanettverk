package collect

import (
	"sort"
)

// keys returns the keys of the map sorted in ascending order.
func keys(m map[string]int) []string {
	k := make([]string, 0, len(m))
	for key := range m {
		k = append(k, key)
	}
	sort.Strings(k) // Sort the keys
	return k
}

// values returns the values of the map sorted in ascending order.
func values(m map[string]int) []int {
	v := make([]int, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}
	sort.Ints(v) // Sort the values
	return v
}
