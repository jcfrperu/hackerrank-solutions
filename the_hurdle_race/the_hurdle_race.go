package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	k := SplitGetIntAt(lines[0], " ", 1)
	heights := SplitInts(lines[1], " ")

	maxHeight := heights[0]
	for _, height := range heights {
		if height > maxHeight {
			maxHeight = height
		}
	}

	result := max(0, maxHeight-k)
	fmt.Fprintf(writer, "%d", result)
}

func main() {
	RunWithFile(solution, "the_hurdle_race/testcases/001-in.txt")
	//RunWithString(solution, "5 4\n1 6 3 5 2")
}

// https://www.hackerrank.com/challenges/the-hurdle-race/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
