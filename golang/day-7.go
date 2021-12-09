package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Define current day for input
const day = 7

func solution1(input string) int64 {
	num_s := strings.Split(strings.TrimSpace(input), ",")

	numbers := []int{}
	for _, s := range num_s {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	sort.Ints(numbers)
	median := numbers[len(numbers)/2]

	var cost int64 = 0
	for _, n := range numbers {
		cost += int64(math.Abs(float64(median - n)))
	}

	return cost
}

func solution2(input string) int64 {
	num_s := strings.Split(strings.TrimSpace(input), ",")

	numbers := []int{}
	for _, s := range num_s {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	min, max := math.MaxInt64, math.MinInt64
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	var minSum int64 = math.MaxInt64
	for i := min; i <= max; i++ {
		var sum int64 = 0
		for _, n := range numbers {
			steps := math.Abs(float64(i - n))
			sum += int64(steps * (steps + 1) / 2)
		}
		if sum < minSum {
			minSum = sum
		}
	}

	return minSum
}

var test_input string = `
16,1,2,0,4,2,7,1,2,14
`

var test_output = map[string]int64{
	"part1": 37,
	"part2": 168,
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
