package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(contents), "\n")

	part1Answer := part1(rows)
	log.Println("Answer to part 1: ", part1Answer)

	part2Answer := part2(rows)
	log.Println("Answer to part 2: ", part2Answer)
}

func part1(input []string) int {
	ventCounts := make(map[pos]int)
	for _, row := range input {
		line := rowToLine(row)

		if line.start.x == line.end.x {
			addVertical(ventCounts, line)
		}
		if line.start.y == line.end.y {
			addHorizontal(ventCounts, line)
		}
	}

	dangerCount := 0
	for _, count := range ventCounts {
		if count >= 2 {
			dangerCount++
		}
	}
	return dangerCount
}

func part2(input []string) int {
	ventCounts := make(map[pos]int)
	for _, row := range input {
		line := rowToLine(row)

		if line.start.x == line.end.x {
			addVertical(ventCounts, line)
			continue
		}
		if line.start.y == line.end.y {
			addHorizontal(ventCounts, line)
			continue
		}

		addDiagonal(ventCounts, line)
	}

	dangerCount := 0
	for _, count := range ventCounts {
		if count >= 2 {
			dangerCount++
		}
	}
	return dangerCount
}

func addVertical(ventCounts map[pos]int, l line) {
	if l.start.y > l.end.y {
		l.start.y, l.end.y = l.end.y, l.start.y
	}
	for y := l.start.y; y <= l.end.y; y++ {
		p := pos{
			x: l.start.x,
			y: y,
		}
		ventCounts[p] = ventCounts[p] + 1
	}
}

func addHorizontal(ventCounts map[pos]int, l line) {
	if l.start.x > l.end.x {
		l.start.x, l.end.x = l.end.x, l.start.x
	}
	for x := l.start.x; x <= l.end.x; x++ {
		p := pos{
			x: x,
			y: l.start.y,
		}
		ventCounts[p] = ventCounts[p] + 1
	}
}

func addDiagonal(ventCounts map[pos]int, l line) {
	if l.start.x > l.end.x {
		l.start.x, l.end.x = l.end.x, l.start.x
		l.start.y, l.end.y = l.end.y, l.start.y
	}
	lineLen := l.end.x - l.start.x

	yInc := 1
	if l.start.y > l.end.y {
		yInc = -1
	}
	for i := 0; i <= lineLen; i++ {
		p := pos{
			x: l.start.x + i,
			y: l.start.y + (i * yInc),
		}
		ventCounts[p] = ventCounts[p] + 1
	}
}

type line struct {
	start, end pos
}

type pos struct {
	x, y int
}

func rowToLine(row string) line {
	parts := strings.Fields(row)
	if len(parts) != 3 {
		log.Fatalf("Expected 3 parts to the string, got %d", len(parts))
	}
	startStrings := strings.Split(parts[0], ",")
	if len(startStrings) != 2 {
		log.Fatalf("Expected 2 parts to the start string, got %d", len(startStrings))
	}
	endStrings := strings.Split(parts[2], ",")
	if len(endStrings) != 2 {
		log.Fatalf("Expected 2 parts to the end string, got %d", len(endStrings))
	}

	// x1,y1 -> x2,y2
	x1, err := strconv.Atoi(startStrings[0])
	if err != nil {
		log.Fatal(err)
	}
	y1, err := strconv.Atoi(startStrings[1])
	if err != nil {
		log.Fatal(err)
	}

	x2, err := strconv.Atoi(endStrings[0])
	if err != nil {
		log.Fatal(err)
	}
	y2, err := strconv.Atoi(endStrings[1])
	if err != nil {
		log.Fatal(err)
	}

	return line{
		start: pos{
			x: x1,
			y: y1,
		},
		end: pos{
			x: x2,
			y: y2,
		},
	}
}
