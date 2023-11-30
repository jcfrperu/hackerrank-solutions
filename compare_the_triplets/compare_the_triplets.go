package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {

	result := []int32{0, 0}
	for i, _ := range a {

		if a[i] > b[i] {
			result[0] += 1
		} else if b[i] > a[i] {
			result[1] += 1
		}
	}

	return result
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
