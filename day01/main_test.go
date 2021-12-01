package main

import "testing"

var input = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func Test_part1(t *testing.T) {
	actual := part1(input)
	if actual != 7 {
		t.Errorf("Wanted 7, got %d", actual)
	}
}

func Test_part2(t *testing.T) {
	actual := part2(input)
	if actual != 5 {
		t.Errorf("Wanted 5, got %d", actual)
	}
}
