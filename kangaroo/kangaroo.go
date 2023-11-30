package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the kangaroo function below.
func kangaroo(x1 float64, v1 float64, x2 float64, v2 float64) string {

	// problem's condition: v2 > v1
	if v2 >= v1 {
		return "NO"
	}

	// position must be integer
	position := (x2 - x1) / (v1 - v2)
	positionInt := float64(int(position))

	// if has decimal part
	if position-positionInt > 0 {
		return "NO"
	}

	return "YES"
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
