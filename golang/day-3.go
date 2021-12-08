package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Define current day for input
const day = 3

func solution1(input string) int64 {
	binary_list := strings.Split(strings.TrimSpace(input), "\n")
	gamma_s, epsilon_s := "", ""
	for i := 0; i < len(binary_list[0]); i++ {
		count := map[string]int{"0": 0, "1": 0}
		for _, binary := range binary_list {
			count[string(binary[i])]++
		}

		if count["0"] > count["1"] {
			gamma_s += "0"
			epsilon_s += "1"
			continue
		}
		gamma_s += "1"
		epsilon_s += "0"
	}

	gamma, _ := strconv.ParseInt(gamma_s, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilon_s, 2, 64)
	return int64(gamma * epsilon)

}

func solution2(input string) int64 {
	binary_list := strings.Split(strings.TrimSpace(input), "\n")
	oxygen_list := make([]string, len(binary_list))
	co2_list := make([]string, len(binary_list))
	copy(oxygen_list, binary_list)
	copy(co2_list, binary_list)

	for i := 0; i < len(oxygen_list[0]) && len(oxygen_list) > 1; i++ {
		count := map[string]int{"0": 0, "1": 0}
		for _, binary := range oxygen_list {
			count[string(binary[i])]++
		}

		cmp := ""
		if count["0"] > count["1"] {
			cmp = "0"
		} else if count["0"] < count["1"] {
			cmp = "1"
		} else {
			cmp = "1"
		}

		oxygen_tmp := []string{}
		for _, s := range oxygen_list {
			if string(s[i]) == cmp {
				oxygen_tmp = append(oxygen_tmp, s)
			}
		}
		oxygen_list = oxygen_tmp
	}

	for i := 0; i < len(co2_list[0]) && len(co2_list) > 1; i++ {
		count := map[string]int{"0": 0, "1": 0}
		for _, binary := range co2_list {
			count[string(binary[i])]++
		}

		cmp := ""
		if count["0"] > count["1"] {
			cmp = "1"
		} else if count["0"] < count["1"] {
			cmp = "0"
		} else {
			cmp = "0"
		}

		co2_tmp := []string{}
		for _, s := range co2_list {
			if string(s[i]) == cmp {
				co2_tmp = append(co2_tmp, s)
			}
		}
		co2_list = co2_tmp
	}

	oxygen, _ := strconv.ParseInt(oxygen_list[0], 2, 64)
	co2, _ := strconv.ParseInt(co2_list[0], 2, 64)
	return int64(oxygen * co2)
}

var test_input string = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

var test_output = map[string]int64{
	"part1": 198,
	"part2": 230,
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
