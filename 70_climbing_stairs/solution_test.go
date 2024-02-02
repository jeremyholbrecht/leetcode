package main

import "testing"

func TestClimbStairs(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		got := climbStairs(3)
		want := 3
		assertCorretAmountOfDistinctWays(t, got, want)
	})
}

func assertCorretAmountOfDistinctWays(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v", "want %v", got, want)
	}
}
