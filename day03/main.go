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
	rows := strings.Split(string(content), "\n")
	uintInput, err := stringsToUint32(rows)
	if err != nil {
		log.Fatalf("Failed to parse strings: %s", err)
	}

	part1Answer := part1(uintInput, len(rows[0]))
	log.Printf("Part 1 answer: %d", part1Answer)
}

func part1(input []uint32, lengthInit int) int {
	p := part1Power(input, lengthInit)
	return int(p.epsilon) * int(p.gamma)
}

type power struct {
	gamma, epsilon uint32
}

func part1Power(input []uint32, lengthInit int) power {
	gamaCounts := make([]int, lengthInit)
	epsilonCounts := make([]int, lengthInit)
	for _, in := range input {
		for i := 0; i < lengthInit; i++ {
			if (in & (1 << i)) != 0 {
				gamaCounts[i]++
			} else {
				epsilonCounts[i]++
			}
		}
	}

	var p power
	for i := 0; i < lengthInit; i++ {
		if gamaCounts[i] > epsilonCounts[i] {
			p.gamma |= (1 << i)
		} else {
			p.epsilon |= (1 << i)
		}
	}
	return p
}

func stringsToUint32(strings []string) ([]uint32, error) {
	uints := make([]uint32, 0, len(strings))
	for _, s := range strings {
		u, err := strconv.ParseUint(s, 2, 32)
		if err != nil {
			return nil, err
		}
		uints = append(uints, uint32(u))
	}
	return uints, nil
}
