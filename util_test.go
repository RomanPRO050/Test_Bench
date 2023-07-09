package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := map[string][]int{
		"1": {1, 2, 3, 4, 5, 1},
		"2": {7, 8, 9, 4, 9, 9, 1},
		"3": {74, 8, 6, 2, 7},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := ContainsDuplicate(testCase)
			require.Exactly(t, res, true)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		IsPalindrome(121)
	})
	t.Run("UnSuccess", func(t *testing.T) {
		IsPalindrome(123)
	})
	t.Run("Guess What?", func(t *testing.T) {
		IsPalindrome(111)

	})
	t.Run("Guess What?", func(t *testing.T) {
		res := IsPalindrome(-121)
		require.Exactly(t, false, res)
	})
}
