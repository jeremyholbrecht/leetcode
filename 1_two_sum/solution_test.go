package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("Example 1:", func(t *testing.T) {
		got := twoSum([]int{2, 7, 11, 15}, 9)
		want := []int{0, 1}
		assertCorrectIndices(t, got, want)

	})

	t.Run("Example 2:", func(t *testing.T) {
		got := twoSum([]int{3, 2, 4}, 6)
		want := []int{1, 2}
		assertCorrectIndices(t, got, want)

	})

}

func assertCorrectIndices(t testing.TB, got []int, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
