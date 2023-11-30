package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here

	var diag01 int32 = 0
	var diag02 int32 = 0

	for i, row := range arr {
		for j, _ := range row {

			if i == j {
				diag01 += arr[i][j]
			}

			if j == len(arr)-i-1 {
				//fmt.Printf(" i=%d j=%d v=%d \n", i, j, arr[i][j])
				diag02 += arr[i][j]
			}
		}
	}
	return int32(math.Abs(float64(diag01 - diag02)))
}

func solution(lines []string, writer *bufio.Writer) {

	Out(writer, fmt.Sprintf("%d", 4))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
