package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Define current day for input
const day = 1

func solution1(input string) int64 {
	numbers_s := strings.Split(strings.TrimSpace(input), "\n")
	numbers := []int{}
	for _, s := range numbers_s {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	count := 0
	for i, n := range numbers {
		if i != 0 && n > numbers[i-1] {
			count++
		}
	}

	return int64(count)
}

func solution2(input string) int64 {
	numbers_s := strings.Split(strings.TrimSpace(input), "\n")
	numbers := []int{}
	for _, s := range numbers_s {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	count := 0
	for i := range numbers {
		if i > 2 {
			prev := numbers[i-1] + numbers[i-2] + numbers[i-3]
			curr := numbers[i] + numbers[i-1] + numbers[i-2]
			if curr > prev {
				count++
			}
		}
	}

	return int64(count)
}

var test_input string = `
199
200
208
210
200
207
240
269
260
263
`

var test_output = map[string]int64{
	"part1": 7,
	"part2": 5,
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
