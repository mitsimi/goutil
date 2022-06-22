package goutil

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestReverse(t *testing.T) {
	gotByte := []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	wantByte := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("reverse byte slice", func(t *testing.T) {
		Reverse(&gotByte)
		for i, v := range gotByte {
			if v != wantByte[i] {
				t.Errorf("got %v, want %v", gotByte, wantByte)
			}
		}
	})

	gotString := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	wantString := []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}

	t.Run("reverse string slice", func(t *testing.T) {
		Reverse(&gotString)
		for i, v := range gotString {
			if v != wantString[i] {
				t.Errorf("got %v, want %v", gotString, wantString)
			}
		}
	})
}

func TestShuffle(t *testing.T) {
	var got []int
	for i := 0; i < 100; i++ {
		got = append(got, i)
	}

	var trys int = 1000
	var fails int = 0

	for i := 0; i < trys; i++ {
		t.Run("shuffle integers", func(t *testing.T) {
			Shuffle(&got)
			var count int = 0
			for i, v := range got {
				if v == i+1 {
					if count >= 3 {
						fails++
					}
					count++
				}
			}
		})
	}
	if condition := fails > trys/10; condition {
		t.Errorf("%v shuffle fails out of %v trys", fails, trys)
	}
}

func TestSorterAsc(t *testing.T) {
	gotInt := []int32{
		-37, -88, -94, 5, -37, 33, 8, -2, 2, 63,
		-43, -96, -98, 5, 0, 40, 59, -38, -21, 33,
		57, -81, 36, 31, 33, 5, -94, 37, 6, -66,
		-31, -26, 11, 40, 58, 3, -41, -25, -74, 28,
		-62, 0, -78, 10, 29, 3, -41, -7, 63, 2,
	}
	wantInt := []int32{
		-98, -96, -94, -94, -88, -81, -78, -74, -66, -62,
		-43, -41, -41, -38, -37, -37, -31, -26, -25, -21,
		-7, -2, 0, 0, 2, 2, 3, 3, 5, 5,
		5, 6, 8, 10, 11, 28, 29, 31, 33, 33,
		33, 36, 37, 40, 40, 57, 58, 59, 63, 63,
	}

	t.Run("Sorting int32 slice", func(t *testing.T) {
		// Sort the int slice.
		SortAsc(&gotInt)
		for i, v := range gotInt {
			if v != wantInt[i] {
				t.Errorf("got %v, want %v", gotInt, wantInt)
			}
		}
	})

	gotByte := []byte{5, 8, 6, 4, 3, 7, 2, 1, 10, 9}
	wantByte := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("Sorting byte slice", func(t *testing.T) {
		// Sort the int slice.
		SortAsc(&gotByte)
		for i, v := range gotByte {
			if v != wantByte[i] {
				t.Errorf("got %v, want %v", gotByte, wantByte)
			}
		}
	})

	t.Run("Sorting string slice", func(t *testing.T) {
		// Generate a slice with random strings.
		var stringSlice []string
		for i := 0; i < 10; i++ {
			stringSlice = append(stringSlice, strconv.Itoa(rand.Intn(100)))
		}

		// Sort the string slice.
		SortAsc(&stringSlice)
		for i, v := range stringSlice {
			if v != stringSlice[i] {
				t.Errorf("string slice %v is not sorted", stringSlice)
				break
			}
		}
	})
}

func TestSorterDesc(t *testing.T) {
	gotInt := []int32{
		-37, -88, -94, 5, -37, 33, 8, -2, 2, 63,
		-43, -96, -98, 5, 0, 40, 59, -38, -21, 33,
		57, -81, 36, 31, 33, 5, -94, 37, 6, -66,
		-31, -26, 11, 40, 58, 3, -41, -25, -74, 28,
		-62, 0, -78, 10, 29, 3, -41, -7, 63, 2,
	}
	wantInt := []int32{
		63, 63, 59, 58, 57, 40, 40, 37, 36, 33,
		33, 33, 31, 29, 28, 11, 10, 8, 6, 5,
		5, 5, 3, 3, 2, 2, 0, 0, -2, -7,
		-21, -25, -26, -31, -37, -37, -38, -41, -41, -43,
		-62, -66, -74, -78, -81, -88, -94, -94, -96, -98,
	}

	t.Run("Sorting int32 slice", func(t *testing.T) {
		// Sort the int slice.
		SortDesc(&gotInt)
		for i, v := range gotInt {
			if v != wantInt[i] {
				t.Errorf("got %v, want %v", gotInt, wantInt)
			}
		}
	})

	gotByte := []byte{5, 8, 6, 4, 3, 7, 2, 1, 10, 9}
	wantByte := []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	t.Run("Sorting byte slice", func(t *testing.T) {
		// Sort the int slice.
		SortDesc(&gotByte)
		for i, v := range gotByte {
			if v != wantByte[i] {
				t.Errorf("got %v, want %v", gotByte, wantByte)
			}
		}
	})

	t.Run("Sorting string slice", func(t *testing.T) {
		// Generate a slice with random strings.
		var stringSlice []string
		for i := 0; i < 10; i++ {
			stringSlice = append(stringSlice, strconv.Itoa(rand.Intn(100)))
		}

		// Sort the string slice.
		SortDesc(&stringSlice)
		for i, v := range stringSlice {
			if v != stringSlice[i] {
				t.Errorf("string slice %v is not sorted", stringSlice)
				break
			}
		}
	})
}

func TestCopy(t *testing.T) {
	got := []float32{
		-37.0, -88.11113, -94.75, 5.676869, -37.0, 33.1, 8.2, -2.9, 2.67, 63.24,
	}
	want := []float32{
		-37.0, -88.11113, -94.75, 5.676869, -37.0, 33.1, 8.2, -2.9, 2.67, 63.24,
	}

	t.Run("Copy", func(t *testing.T) {
		// Copy the float slice.
		cp, err := Copy(&got)
		if err != nil {
			t.Errorf("got error %v, want no error", err)
		}

		// Check if the slices are equal.
		if !reflect.DeepEqual(cp, want) {
			t.Errorf("got %v, want %v", cp, want)
		}

	})

}
