package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	list := SplitInts(lines[1], " ")

	minCount, minVal := 0, list[0]
	maxCount, maxVal := 0, list[0]

	for _, item := range list {
		if item > maxVal {
			maxVal = item
			maxCount++
		}
		if item < minVal {
			minVal = item
			minCount++
		}
	}

	fmt.Fprintf(writer, "%d %d", maxCount, minCount)
}

func main() {
	RunWithFile(solution, "breaking_the_records/testcases/001-in.txt")
	//RunWithString(solution, "9\n10 5 20 20 4 5 2 25 1")
}

// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
