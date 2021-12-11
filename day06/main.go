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

	part1Answer := part1(ints, 80)
	log.Println("Answer for part 1: ", part1Answer)

	part2Answer := part2(ints, 256)
	log.Println("Answer for part 2: ", part2Answer)
}

func part1(in []int, days int) int {
	lanternfish := make([]int, len(in))
	copy(lanternfish, in)
	for d := 0; d < days; d++ {
		newFish := []int{}
		for i := range lanternfish {
			if lanternfish[i] == 0 {
				newFish = append(newFish, 8)
				lanternfish[i] = 6
			} else {
				lanternfish[i]--
			}
		}
		lanternfish = append(lanternfish, newFish...)
	}

	return len(lanternfish)
}

func part2(in []int, days int) int {
	fishes := make([]int, 9)
	for _, i := range in {
		fishes[i]++
	}

	for d := 0; d < days; d++ {
		fishes = append(fishes[1:], fishes[0])
		fishes[6] += fishes[8]
	}
	fishTot := 0
	for _, count := range fishes {
		fishTot += count
	}
	return fishTot
}
