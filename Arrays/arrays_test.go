package main

import "testing"

// for getting coverage of test use: go test -cover

func TestSum(t *testing.T) {
	t.Run("array testing", func(t *testing.T) {
		array := [5]int{1, 2, 3, 4, 5}

		got := Sum(array[:])
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, array)
		}
	})

	t.Run("slice testing", func(t *testing.T) {
		mySlice := []int{1, 2, 3, 4, 5}

		got := Sum(mySlice)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, mySlice)
		}
	})
}
