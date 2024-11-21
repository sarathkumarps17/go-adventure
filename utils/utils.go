package utils

import "reflect"

// CheckSlicesHaveSameElements checks if two slices have the same elements.
// The slices are not required to have the same order, but they must have the same
// elements with the same frequency.
func CheckSlicesHaveSameElements[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	freq1 := make(map[T]int)
	freq2 := make(map[T]int)

	for _, elem := range slice1 {
		freq1[elem]++
	}
	for _, elem := range slice2 {
		freq2[elem]++
	}

	return reflect.DeepEqual(freq1, freq2)
}
