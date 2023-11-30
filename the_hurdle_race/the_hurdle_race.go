package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	k := SplitIntsGetAt(lines[0], " ", 1)
	heights := SplitInts(lines[1], " ")

	max := heights[0]
	for _, height := range heights {
		if height > max {
			max = height
		}
	}

	result := ExecInt("max", 0, max-k)

	Out(writer, fmt.Sprintf("%d", result))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
