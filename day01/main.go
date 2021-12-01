package main

import (
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
	values := strings.Split(string(content), "\n")

	var intVals []int
	for _, v := range values {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %s", err)
		}
		intVals = append(intVals, intVal)
	}

	part1Answer := part1(intVals)
	log.Printf("Answer for part 1: %d\n", part1Answer)

	part2Answer := part2(intVals)
	log.Printf("Answer for part 2: %d", part2Answer)
}

func part1(in []int) int {
	var countIncreases int
	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] {
			countIncreases++
		}
	}
	return countIncreases
}

func part2(in []int) int {
	newSlice := make([]int, 0, len(in))
	for i := 0; i <= len(in)-3; i++ {
		newSlice = append(newSlice, sum(in[i:i+3]))
	}
	return part1(newSlice)
}

func sum(in []int) int {
	var tot int
	for _, i := range in {
		tot += i
	}
	return tot
}
