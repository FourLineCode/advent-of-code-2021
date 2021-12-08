package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Define current day for input
const day = 4

func solution1(input string) int64 {
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	numbers_block, boards_block := blocks[0], blocks[1:]

	numbers_s := strings.Split(numbers_block, ",")
	numbers := []int{}
	for _, s := range numbers_s {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, num)
	}

	boards := [][][]int{}
	visited := [][][]bool{}
	for _, block := range boards_block {
		board := [][]int{}
		vis := [][]bool{}
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			n_s := strings.Split(line, " ")
			n_list := []int{}
			vis_list := []bool{}
			for _, n := range n_s {
				if n != "" {
					num, _ := strconv.Atoi(strings.TrimSpace(n))
					n_list = append(n_list, num)
					vis_list = append(vis_list, false)
				}
			}
			board = append(board, n_list)
			vis = append(vis, vis_list)
		}
		boards = append(boards, board)
		visited = append(visited, vis)
	}

	for _, num := range numbers {
		for b_i, board := range boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board[i][j] == num {
						visited[b_i][i][j] = true
					}
				}
			}

			for i := 0; i < 5; i++ {
				won_row, won_col := true, true
				for j := 0; j < 5; j++ {
					if visited[b_i][i][j] == false {
						won_row = false
						break
					}
				}
				for j := 0; j < 5; j++ {
					if visited[b_i][j][i] == false {
						won_col = false
						break
					}
				}
				if won_row || won_col {
					sum := 0
					for i := 0; i < 5; i++ {
						for j := 0; j < 5; j++ {
							if visited[b_i][i][j] == false {
								sum += board[i][j]
							}
						}
					}
					return int64(sum * num)
				}
			}
		}
	}

	return -1
}

func solution2(input string) int64 {
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	numbers_block, boards_block := blocks[0], blocks[1:]

	numbers_s := strings.Split(numbers_block, ",")
	numbers := []int{}
	for _, s := range numbers_s {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, num)
	}

	boards := [][][]int{}
	visited := [][][]bool{}
	for _, block := range boards_block {
		board := [][]int{}
		vis := [][]bool{}
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			n_s := strings.Split(line, " ")
			n_list := []int{}
			vis_list := []bool{}
			for _, n := range n_s {
				if n != "" {
					num, _ := strconv.Atoi(strings.TrimSpace(n))
					n_list = append(n_list, num)
					vis_list = append(vis_list, false)
				}
			}
			board = append(board, n_list)
			vis = append(vis, vis_list)
		}
		boards = append(boards, board)
		visited = append(visited, vis)
	}

	won_index := []int{}
	for _, num := range numbers {
		for b_i, board := range boards {
			already_won := false
			for _, index := range won_index {
				if index == b_i {
					already_won = true
					break
				}
			}
			if already_won {
				continue
			}

			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board[i][j] == num {
						visited[b_i][i][j] = true
					}
				}
			}

			for i := 0; i < 5; i++ {
				won_row, won_col := true, true
				for j := 0; j < 5; j++ {
					if visited[b_i][i][j] == false {
						won_row = false
						break
					}
				}
				for j := 0; j < 5; j++ {
					if visited[b_i][j][i] == false {
						won_col = false
						break
					}
				}
				won := won_row || won_col
				if won {
					if len(boards)-len(won_index) == 1 {
						sum := 0
						for i := 0; i < 5; i++ {
							for j := 0; j < 5; j++ {
								if visited[b_i][i][j] == false {
									sum += board[i][j]
								}
							}
						}
						return int64(sum * num)
					} else {
						exists := false
						for _, index := range won_index {
							if index == b_i {
								exists = true
							}
						}
						if !exists {
							won_index = append(won_index, b_i)
						}
					}
				}
			}
		}
	}

	return -1
}

var test_input string = `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`

var test_output = map[string]int64{
	"part1": 4512,
	"part2": 1924,
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
