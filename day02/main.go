package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	rows := strings.Split(string(content), "\n")
	commands, err := stringsToCommands(rows)
	if err != nil {
		log.Fatalf("Failed to convert input rows into commands: %s", err)
	}

	part1Answer := part1(commands)
	log.Printf("Answer for part 1: %d\n", part1Answer)

	part2Answer := part2(commands)
	log.Printf("Answer for part 2: %d\n", part2Answer)
}

type command struct {
	direction string
	amplitude int
}

type position struct {
	horizontal, depth int
}

func part1(input []command) int {
	pos := calcPosition(input)
	return pos.depth * pos.horizontal
}

func part2(input []command) int {
	pos := calcPositionPart2(input)
	return pos.depth * pos.horizontal
}

func calcPosition(input []command) position {
	var pos position
	for _, c := range input {
		switch c.direction {
		case "forward":
			pos.horizontal += c.amplitude
		case "up":
			pos.depth -= c.amplitude
		case "down":
			pos.depth += c.amplitude
		}
	}
	return pos
}

func calcPositionPart2(input []command) position {
	var pos position
	var aim int
	for _, c := range input {
		switch c.direction {
		case "forward":
			pos.horizontal += c.amplitude
			pos.depth += aim * c.amplitude
		case "up":
			aim -= c.amplitude
		case "down":
			aim += c.amplitude
		}
	}
	return pos
}

func stringsToCommands(s []string) ([]command, error) {
	coms := make([]command, 0, len(s))
	for _, row := range s {
		r := strings.Split(row, " ")
		if len(r) != 2 {
			return nil, fmt.Errorf("expected two items, got %d", len(r))
		}
		amp, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, err
		}
		coms = append(coms, command{
			direction: r[0],
			amplitude: amp,
		})
	}
	return coms, nil
}
