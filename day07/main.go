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
		log.Fatal(err)
	}
	vals := strings.Split(string(content), ",")
	ints := make([]int, len(vals))
	for i, s := range vals {
		ints[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
	}

	part1Answer := part1(ints)
	log.Println("Answer for part 1: ", part1Answer)
}

func part1(in []int) int {
	min, max := minMax(in)

	minFuel := 10000000 // Some large number
	for i := min; i <= max; i++ {
		fuel := fuelUsed(in, i)
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func minMax(in []int) (int, int) {
	min := 1000
	max := 0
	for i := range in {
		if in[i] > max {
			max = in[i]
		}
		if in[i] < min {
			min = in[i]
		}
	}

	return min, max
}

func fuelUsed(in []int, moveTo int) int {
	fuel := 0
	for i := range in {
		d := absDiff(in[i], moveTo)
		fuel += d
	}
	return fuel
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
