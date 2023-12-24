package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	list := SplitInts(lines[0], " ")
	values := lines[1]

	maxValue := -1
	for _, r := range values {
		i := r - 'a'
		if list[i] > maxValue {
			maxValue = list[i]
		}
	}

	result := maxValue * len(values)

	fmt.Fprintf(writer, "%d", result)
}

func main() {
	RunWithFile(solution, "designer-pfd-viewer/testcases/001-in.txt")
	//RunWithString(solution, "1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5\nabc")
}

// https://www.hackerrank.com/challenges/designer-pdf-viewer/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
