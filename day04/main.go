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

	calledStrings := strings.Split(rows[0], ",")
	calledNums := make([]int, 0, len(calledStrings))
	for _, s := range calledStrings {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Failed to convert called number from string to int: %v", err)
		}
		calledNums = append(calledNums, i)
	}

	part1Answer := part1(calledNums, rows[2:])
	log.Printf("Answer to part 1: %d", part1Answer)

}

func part1(calledNums []int, rows []string) int {
	var boardRows []string
	for _, r := range rows {
		if r != "" {
			boardRows = append(boardRows, r)
		}
	}

	var boardStrings [][]string
	for i := 0; i < len(boardRows)-1; i += 5 {
		boardStrings = append(boardStrings, boardRows[i:i+5])
	}

	var boards []board
	for _, b := range boardStrings {
		boards = append(boards, stringsToBoard(b))
	}

	return part1Bingo(boards, calledNums)
}

type board struct {
	rows    [][]boardNumber
	columns [][]boardNumber
}

type boardNumber struct {
	value     int
	completed bool
}

func stringsToBoard(rows []string) board {
	var b board
	for _, r := range rows {
		numStrings := strings.Fields(r)

		var row []boardNumber
		for _, c := range numStrings {
			num, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("Failed to parse col string to int: %v", err)
			}
			boardNum := boardNumber{
				value:     num,
				completed: false,
			}
			row = append(row, boardNum)
		}

		b.rows = append(b.rows, row)
	}

	for i := 0; i < 5; i++ {
		var col []boardNumber
		for _, row := range b.rows {
			col = append(col, row[i])
		}
		b.columns = append(b.columns, col)
	}

	return b
}

func part1Bingo(boards []board, calledNumbers []int) int {
	for _, n := range calledNumbers {
		for _, b := range boards {
			for k, r := range b.rows {
				var completedRowCount int
				for j, i := range r {
					if i.value == n {
						i.completed = true
					}
					if i.completed {
						completedRowCount++
					}
					r[j] = i
				}
				b.rows[k] = r
				if completedRowCount == 5 {
					winningNum := n
					sumNotCompleted := sumUnmarked(b)
					return winningNum * sumNotCompleted
				}
			}

			for k, c := range b.columns {
				var completedColCount int
				for j, i := range c {
					if i.value == n {
						i.completed = true
					}
					if i.completed {
						completedColCount++
					}
					c[j] = i
				}
				b.columns[k] = c
				if completedColCount == 5 {
					winningNum := n
					sumNotCompleted := sumUnmarked(b)
					return winningNum * sumNotCompleted
				}
			}
		}
	}

	return 0
}

func sumUnmarked(b board) int {
	var unmarked int
	for _, r := range b.rows {
		for _, i := range r {
			if !i.completed {
				unmarked += i.value
			}
		}
	}
	return unmarked
}
