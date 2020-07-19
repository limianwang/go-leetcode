package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	tt := []struct {
		name        string
		inputCoins  []int
		inputAmount int
		expected    int
	}{
		{
			name:        "test 1",
			inputCoins:  []int{2},
			inputAmount: 3,
			expected:    -1,
		},
		{
			name:        "test 2",
			inputCoins:  []int{1, 2, 5},
			inputAmount: 11,
			expected:    3,
		},
		{
			name:        "test 3",
			inputCoins:  []int{1},
			inputAmount: 0,
			expected:    0,
		},
		{
			name:        "test 4",
			inputCoins:  []int{1, 2, 5},
			inputAmount: 100,
			expected:    20,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if actual := coinChange(tc.inputCoins, tc.inputAmount); actual != tc.expected {
				t.Fatalf(
					"error getting valid coin amount wanted %d but got %d",
					tc.expected,
					actual,
				)
			}
		})
	}
}
