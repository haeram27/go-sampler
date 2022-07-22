package util

// slice substraction: O(n)
// remove v elements from slice
func RemoveFromSlice[T comparable](l []T, v T) []T {
	for i, e := range l {
		if e == v {
			l = append(l[:i], l[i+1:]...)
		}
	}

	return l
}

// slice substraction:  O(1)
// remove elements of slice b from slice a return a - b
func DiffStrSlice(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, v := range b {
		mb[v] = struct{}{}
	}

	var diff []string
	for _, v := range a {
		if _, found := mb[v]; !found {
			diff = append(diff, v)
		}
	}

	return diff
}
