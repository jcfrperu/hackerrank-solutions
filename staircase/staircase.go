package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the staircase function below.
func staircase(n int32) {

	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n; j++ {
			if j >= n-i-1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		if i != n-1 {
			fmt.Printf("\n")
		}
	}
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
