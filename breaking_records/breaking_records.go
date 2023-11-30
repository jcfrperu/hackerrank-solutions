package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	list := SplitInts(lines[1], " ")

	min := list[0]
	max := list[0]

	minCount := 0
	maxCount := 0

	for _, item := range list {
		if item > max {
			max = item
			maxCount++
		}
		if item < min {
			min = item
			minCount++
		}
	}

	Out(writer, fmt.Sprintf("%d %d", maxCount, minCount))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
