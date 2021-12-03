package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := []command{
		{direction: "forward", amplitude: 5},
		{direction: "down", amplitude: 5},
		{direction: "forward", amplitude: 8},
		{direction: "up", amplitude: 3},
		{direction: "down", amplitude: 8},
		{direction: "forward", amplitude: 2},
	}

	actual := part1(input)

	if actual != 150 {
		t.Errorf("Expected 150, got %d", actual)
	}
}

func Test_part2(t *testing.T) {
	input := []command{
		{direction: "forward", amplitude: 5},
		{direction: "down", amplitude: 5},
		{direction: "forward", amplitude: 8},
		{direction: "up", amplitude: 3},
		{direction: "down", amplitude: 8},
		{direction: "forward", amplitude: 2},
	}

	actual := calcPositionPart2(input)

	expected := position{
		horizontal: 15,
		depth:      60,
	}
	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func Test_stringsToCommands(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	expected := []command{
		{direction: "forward", amplitude: 5},
		{direction: "down", amplitude: 5},
		{direction: "forward", amplitude: 8},
		{direction: "up", amplitude: 3},
		{direction: "down", amplitude: 8},
		{direction: "forward", amplitude: 2},
	}

	actual, err := stringsToCommands(input)
	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err)
	}
	if len(actual) != len(expected) {
		t.Errorf("Expected len %d, got len %d", len(expected), len(actual))
	}
	for i, a := range actual {
		if a != expected[i] {
			t.Errorf("Expected %+v, got %+v", expected[i], a)
		}
	}
}
