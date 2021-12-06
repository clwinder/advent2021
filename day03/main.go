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

	part2Answer := part2(rows)
	log.Printf("Part 2 answer: %d", part2Answer)
}

func part1(input []uint32, lengthInit int) int {
	p := part1Power(input, lengthInit)
	return int(p.epsilon) * int(p.gamma)
}

func part2(input []string) int {
	l := part2LifeSupport(input)
	return l.oxygen * l.co2
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

type lifeSupport struct {
	oxygen, co2 int
}

func part2LifeSupport(input []string) lifeSupport {
	return lifeSupport{
		oxygen: oxygen(input),
		co2:    co2(input),
	}
}

func oxygen(input []string) int {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	for i := 0; i < len(inputCopy[0]); i++ {
		var zeroCount, oneCount int
		for _, in := range inputCopy {
			if in[i] == '0' {
				zeroCount++
			}
			if in[i] == '1' {
				oneCount++
			}
		}

		keep := []string{}
		for _, in := range inputCopy {
			if oneCount >= zeroCount {
				if in[i] == '1' {
					keep = append(keep, in)
				}
			} else {
				if in[i] == '0' {
					keep = append(keep, in)
				}
			}
		}
		inputCopy = keep
		if len(inputCopy) == 1 {
			oxygen, err := strconv.ParseInt(inputCopy[0], 2, 64)
			if err != nil {
				log.Fatalf("Failed to convert oxygen string to int: %v", err)
			}
			return int(oxygen)
		}
	}
	return 0
}

func co2(input []string) int {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	for i := 0; i < len(inputCopy[0]); i++ {
		var zeroCount, oneCount int
		for _, in := range inputCopy {
			if in[i] == '0' {
				zeroCount++
			}
			if in[i] == '1' {
				oneCount++
			}
		}

		keep := []string{}
		for _, in := range inputCopy {
			if oneCount >= zeroCount {
				if in[i] == '0' {
					keep = append(keep, in)
				}
			} else {
				if in[i] == '1' {
					keep = append(keep, in)
				}
			}
		}
		inputCopy = keep
		if len(inputCopy) == 1 {
			oxygen, err := strconv.ParseInt(inputCopy[0], 2, 64)
			if err != nil {
				log.Fatalf("Failed to convert oxygen string to int: %v", err)
			}
			return int(oxygen)
		}
	}
	return 0
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
