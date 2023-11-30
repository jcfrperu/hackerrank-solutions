package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	list := SplitInts(lines[0], " ")
	values := lines[1]

	var max = -1
	for _, r := range values {
		i := r - 'a'
		if list[i] > max {
			max = list[i]
		}
	}

	result := max * len(values)

	Out(writer, fmt.Sprintf("%d", result))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
