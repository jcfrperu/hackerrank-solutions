package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	// first line
	var s = SplitIntsGetAt(lines[0], " ", 0)
	var t = SplitIntsGetAt(lines[0], " ", 1)
	// second line
	var a = SplitIntsGetAt(lines[1], " ", 0)
	var b = SplitIntsGetAt(lines[1], " ", 1)
	// third line is ignored
	// fourth line
	var apples = SplitInts(lines[3], " ")
	var oranges = SplitInts(lines[4], " ")

	var applesNumber int64 = 0
	var orangesNumber int64 = 0

	for _, da := range apples {
		if a+da >= s && a+da <= t {
			applesNumber++
		}
	}

	for _, db := range oranges {
		if b+db <= t && b+db >= s {
			orangesNumber++
		}
	}

	Out(writer, fmt.Sprintf("%d\n%d", applesNumber, orangesNumber))
}

func main() {
	//RunWithFile(solution, "apple_and_orange/testcases/003.input")
	RunWithString(solution, "7 11\n5 15\n3 2\n-2 2 1\n5 -6")
}

// https://www.hackerrank.com/challenges/apple-and-orange/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
