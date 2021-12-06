package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	uintInput, err := stringsToUint32(input)
	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err)
	}

	actual := part1(uintInput, len(input[0]))
	if actual != 198 {
		t.Errorf("Expected 198, got %d", actual)
	}
}

func Test_part2LifeSupport(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	actual := part2LifeSupport(input)

	expected := lifeSupport{
		oxygen: 23,
		co2:    10,
	}

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
