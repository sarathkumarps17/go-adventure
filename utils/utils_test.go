package utils

import "testing"

// CheckSlicesHaveSameElements returns true for slices with identical elements in different orders
func TestCheckSlicesHaveSameElementsIdenticalElementsDifferentOrder(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{4, 3, 2, 1}

	result := CheckSlicesHaveSameElements(slice1, slice2)

	if !result {
		t.Errorf("Expected true, got false")
	}
}

// CheckSlicesHaveSameElements returns false for slices of different lengths
func TestCheckSlicesHaveSameElementsDifferentLengths(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3, 4}

	result := CheckSlicesHaveSameElements(slice1, slice2)

	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test slices with strings
func TestCheckSlicesHaveSameElementsWithStrings(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"cherry", "banana", "apple"}

	result := CheckSlicesHaveSameElements(slice1, slice2)

	if !result {
		t.Errorf("Expected true, got false")
	}
}

func TestCheckSlicesHaveSameElementsWithStringsThatAreNotIdentical(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"cherry", "banana", "apple", "orange"}

	result := CheckSlicesHaveSameElements(slice1, slice2)

	if result {
		t.Errorf("Expected False, got True")
	}
}
