package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the birthday function below.
func birthday(s []int, d int, m int) int {

	var matches = 0
	for i := 0; i < len(s)-m+1; i++ {
		var sum = 0
		for j := i; j < i+m; j++ {
			sum += s[j]
		}
		if sum == d {
			matches++
		}
	}

	return matches
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
