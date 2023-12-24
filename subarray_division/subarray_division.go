package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	s := SplitInts(lines[1], " ")
	d := SplitGetIntAt(lines[2], " ", 0)
	m := SplitGetIntAt(lines[2], " ", 1)

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

	fmt.Fprintf(writer, "%d", matches)
}

func main() {
	RunWithFile(solution, "subarray_division/testcases/001-in.txt")
	//RunWithString(solution, "5\n1 2 1 3 2\n3 2")
}

// https://www.hackerrank.com/challenges/the-birthday-bar/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
