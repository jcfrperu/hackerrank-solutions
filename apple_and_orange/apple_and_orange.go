package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	// first line
	s := SplitGetIntAt(lines[0], " ", 0)
	t := SplitGetIntAt(lines[0], " ", 1)
	// second line
	a := SplitGetIntAt(lines[1], " ", 0)
	b := SplitGetIntAt(lines[1], " ", 1)
	// third line is ignored
	// fourth line
	apples := SplitInts(lines[3], " ")
	oranges := SplitInts(lines[4], " ")

	applesNumber := int64(0)
	for _, da := range apples {
		if a+da >= s && a+da <= t {
			applesNumber++
		}
	}

	orangesNumber := int64(0)
	for _, db := range oranges {
		if b+db <= t && b+db >= s {
			orangesNumber++
		}
	}

	fmt.Fprintf(writer, "%d\n%d", applesNumber, orangesNumber)
}

func main() {
	RunWithFile(solution, "apple_and_orange/testcases/001-in.txt")
	//RunWithString(solution, "7 11\n5 15\n3 2\n-2 2 1\n5 -6")
}

// https://www.hackerrank.com/challenges/apple-and-orange/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
