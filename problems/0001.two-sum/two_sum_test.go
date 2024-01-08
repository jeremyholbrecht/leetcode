package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		got := []int{1, 2, 3}
		want := []int{5, 5, 7}

		if slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
