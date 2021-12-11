package main

import "testing"

func Test_part1(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	tests := map[string]struct {
		days     int
		expected int
	}{
		"18 days": {
			days:     18,
			expected: 26,
		},
		"80 days": {
			days:     80,
			expected: 5934,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := part1(input, test.days)
			if actual != test.expected {
				t.Errorf("Expected %d fish, got %d", test.expected, actual)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	tests := map[string]struct {
		days     int
		expected int
	}{
		"18 days": {
			days:     18,
			expected: 26,
		},
		"80 days": {
			days:     80,
			expected: 5934,
		},
		"256 days": {
			days:     256,
			expected: 26984457539,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := part2(input, test.days)
			if actual != test.expected {
				t.Errorf("Expected %d fish, got %d", test.expected, actual)
			}
		})
	}
}
