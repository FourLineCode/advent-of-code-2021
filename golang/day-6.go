package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Define current day for input
const day = 6

func solution1(input string) int64 {
	num_s := strings.Split(strings.TrimSpace(input), ",")

	numbers := []int{}
	for _, s := range num_s {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	bucket := map[int]int64{}
	for _, n := range numbers {
		bucket[n]++
	}

	for i := 0; i < 80; i++ {
		tmp := map[int]int64{}
		for day, count := range bucket {
			if day == 0 {
				tmp[6] += count
				tmp[8] += count
			} else {
				tmp[day-1] += count
			}
		}
		bucket = tmp
	}

	var sum int64 = 0
	for _, count := range bucket {
		sum += count
	}

	return sum
}

func solution2(input string) int64 {
	num_s := strings.Split(strings.TrimSpace(input), ",")

	numbers := []int{}
	for _, s := range num_s {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	bucket := map[int]int64{}
	for _, n := range numbers {
		bucket[n]++
	}

	for i := 0; i < 256; i++ {
		tmp := map[int]int64{}
		for day, count := range bucket {
			if day == 0 {
				tmp[6] += count
				tmp[8] += count
			} else {
				tmp[day-1] += count
			}
		}
		bucket = tmp
	}

	var sum int64 = 0
	for _, count := range bucket {
		sum += count
	}

	return sum
}

var test_input string = `
3,4,3,1,2
`

var test_output = map[string]int64{
	"part1": 5934,
	"part2": 26984457539,
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
