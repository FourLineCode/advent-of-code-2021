package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Define current day for input
const day = 2

func solution1(input string) int64 {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	horizontal, depth := 0, 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		direction := split[0]
		amount, _ := strconv.Atoi(split[1])

		switch direction {
		case "forward":
			horizontal += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}
	return int64(horizontal * depth)
}

func solution2(input string) int64 {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	horizontal, depth, aim := 0, 0, 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		direction := split[0]
		amount, _ := strconv.Atoi(split[1])

		switch direction {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}
	return int64(horizontal * depth)
}

var test_input string = `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`

var test_output = map[string]int64{
	"part1": 150,
	"part2": 900,
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
