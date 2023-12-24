package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	var n = ParseInt(lines[0])
	var p = ParseInt(lines[1])

	var turnsLeft = (p - 0) / 2

	var turnsRight = (n - p) / 2
	if n%2 == 0 && p%2 != 0 {
		turnsRight += 1
	}
	var turns = min(turnsLeft, turnsRight)

	fmt.Fprintf(writer, "%d", turns)
}

func main() {
	//RunWithFile(solution, "drawing_book/testcases/002-in.txt")
	RunWithString(solution, "4\n4")
}

// https://www.hackerrank.com/challenges/drawing-book/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
