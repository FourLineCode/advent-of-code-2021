package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Define current day for input
const day = 5

func solution1(input string) int64 {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	wires := [][][]int{}
	board := [1000][1000]int{}

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		from, to := []int{}, []int{}

		from_s := strings.Split(parts[0], ",")
		for _, s := range from_s {
			n, _ := strconv.Atoi(s)
			from = append(from, n)
		}
		to_s := strings.Split(parts[1], ",")
		for _, s := range to_s {
			n, _ := strconv.Atoi(s)
			to = append(to, n)
		}

		wires = append(wires, [][]int{from, to})
	}

	for _, wire := range wires {
		from, to := wire[0], wire[1]
		x1, y1 := from[0], from[1]
		x2, y2 := to[0], to[1]

		if x1 == x2 {
			for i := math.Min(float64(y1), float64(y2)); i <= math.Max(float64(y1), float64(y2)); i++ {
				board[x1][int(i)]++
			}
		} else if y1 == y2 {
			for i := math.Min(float64(x1), float64(x2)); i <= math.Max(float64(x1), float64(x2)); i++ {
				board[int(i)][y1]++
			}
		}
	}

	count := 0
	for _, row := range board {
		for _, col := range row {
			if col > 1 {
				count++
			}
		}
	}

	return int64(count)
}

func solution2(input string) int64 {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	wires := [][][]int{}
	board := [1000][1000]int{}

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		from, to := []int{}, []int{}

		from_s := strings.Split(parts[0], ",")
		for _, s := range from_s {
			n, _ := strconv.Atoi(s)
			from = append(from, n)
		}
		to_s := strings.Split(parts[1], ",")
		for _, s := range to_s {
			n, _ := strconv.Atoi(s)
			to = append(to, n)
		}

		wires = append(wires, [][]int{from, to})
	}

	for _, wire := range wires {
		from, to := wire[0], wire[1]
		x1, y1 := from[0], from[1]
		x2, y2 := to[0], to[1]

		if x1 == x2 {
			for i := math.Min(float64(y1), float64(y2)); i <= math.Max(float64(y1), float64(y2)); i++ {
				board[x1][int(i)]++
			}
		} else if y1 == y2 {
			for i := math.Min(float64(x1), float64(x2)); i <= math.Max(float64(x1), float64(x2)); i++ {
				board[int(i)][y1]++
			}
		} else {
			xDir, yDir := 1, 1
			if x1 > x2 {
				xDir = -1
			}
			if y1 > y2 {
				yDir = -1
			}

			var i, j int
			for i, j = x1, y1; i != x2 && j != y2; i, j = i+xDir, j+yDir {
				board[i][j]++
			}
			board[i][j]++
		}
	}

	count := 0
	for _, row := range board {
		for _, col := range row {
			if col > 1 {
				count++
			}
		}
	}

	return int64(count)
}

var test_input string = `
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

var test_output = map[string]int64{
	"part1": 5,
	"part2": 12,
}

func main() {
	content, err := ioutil.ReadFile(fmt.Sprintf("../inputs/%v.txt", day))
	if err != nil {
		fmt.Println("Coudn't read input file")
		return
	}
	input := string(content)

	fmt.Println("Running solution 1 ...")
	test1_output := "FAILED"
	test1_color := ColorRed
	if solution1(test_input) == test_output["part1"] {
		test1_output = "PASSED"
		test1_color = ColorGreen
	}
	fmt.Printf("Test for solution 1 - %v%v%v\n", test1_color, test1_output, ColorReset)
	fmt.Printf("Solution 1 - %v%v%v\n\n", ColorYellow, solution1(input), ColorReset)

	fmt.Println("Running solution 1 ...")
	test2_output := "FAILED"
	test2_color := ColorRed
	if solution2(test_input) == test_output["part2"] {
		test2_output = "PASSED"
		test2_color = ColorGreen
	}
	fmt.Printf("Test for solution 2 - %v%v%v\n", test2_color, test2_output, ColorReset)
	fmt.Printf("Solution 2 - %v%v%v\n", ColorYellow, solution2(input), ColorReset)
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)
