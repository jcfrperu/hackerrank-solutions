package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"math"
)

func solution(lines []string, writer *bufio.Writer) {
	arr := make([][]int, 0)
	for _, line := range lines[1:] {
		arr = append(arr, SplitInts(line, " "))
	}

	diag01, diag02 := 0, 0
	for i, row := range arr {
		for j := range row {
			if i == j {
				diag01 += arr[i][j]
			}
			if j == len(arr)-i-1 {
				diag02 += arr[i][j]
			}
		}
	}

	absValue := int(math.Abs(float64(diag01 - diag02)))
	fmt.Fprintf(writer, "%d", absValue)
}

func main() {
	RunWithFile(solution, "diagonal_difference/testcases/001-in.txt")
	//RunWithString(solution, "3\n11 2 4\n4 5 6\n10 8 -12")
}

// https://www.hackerrank.com/challenges/diagonal-difference/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
