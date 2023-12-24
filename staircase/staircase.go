package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	n := ParseInt(lines[0])

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= n-i-1 {
				fmt.Fprintf(writer, "#")
			} else {
				fmt.Fprintf(writer, " ")
			}
		}
		if i != n-1 {
			fmt.Fprintf(writer, "\n")
		}
	}
}

func main() {
	RunWithFile(solution, "staircase/testcases/001-in.txt")
	//RunWithString(solution, "6")
}

// https://www.hackerrank.com/challenges/staircase/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
