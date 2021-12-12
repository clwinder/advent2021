package main

import "testing"

func Test_part1(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	actual := part1(input)
	if actual != 37 {
		t.Errorf("Expected 37, got %d", actual)
	}
}
