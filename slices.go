package goutil

import (
	"errors"
	"math/rand"
)

// Take a slice of any orderable type and sort it in ascending order.
// The type of the slice must be one of the types defined by the type Any, Number.
func SortAsc[A Any](slice *[]A) error {
	if slice == nil {
		return errors.New("nil slice")
	}

	// Range through the slice and sort it.
	for i := 0; i < len(*slice)-1; i++ {
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[i] > (*slice)[j] {
				(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
			}
		}
	}
	return nil
}

//Take a slice of any orderable type and sort it in descending order.
//The type of the slice must be one of the types defined by the type Any, Number.
func SortDesc[A Any](slice *[]A) error {
	if slice == nil {
		return errors.New("nil slice")
	}

	// Range through the slice and sort it.
	for i := 0; i < len(*slice)-1; i++ {
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[i] < (*slice)[j] {
				(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
			}
		}
	}
	return nil
}

// Take a slice of any type and shuffle it.
func Shuffle[A any](slice *[]A) error {
	if slice == nil {
		return errors.New("nil slice")
	}

	// Range through the slice and shuffle it.
	for i := 0; i < len(*slice); i++ {
		j := rand.Intn(len(*slice))
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
	return nil
}

// Take a slice of any type and reverse it.
func Reverse[A any](slice *[]A) error {
	if slice == nil {
		return errors.New("nil slice")
	}

	// Range through the slice and reverse it.
	for i := 0; i < len(*slice)/2; i++ {
		(*slice)[i], (*slice)[len(*slice)-1-i] = (*slice)[len(*slice)-1-i], (*slice)[i]
	}
	return nil
}

// Take a slice of any type and return a new slice with the same elements in the same order.
func CopySlice[A any](slice *[]A) ([]A, error) {
	if slice == nil {
		return nil, errors.New("nil slice")
	}

	// Create a new slice and copy the elements.
	newSlice := make([]A, len(*slice))
	copy(newSlice, *slice)
	return newSlice, nil
}
