package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"io"
)

func solution(lines []string, writer *bufio.Writer) {
	list := SplitInts(lines[1], " ")

	sum := 0
	for i := range list {
		sum += list[i]
	}

	Print2(writer, "%d", sum)
}

func Print2(w io.Writer, format string, a any) {
	fmt.Fprintf(w, format, a)
}

func main() {
	//RunWithFile(solution, "simple_array_sum/testcases/001-in.txt")
	RunWithString(solution, "6\n1 2 3 4 10 11")
}

// https://www.hackerrank.com/challenges/simple-array-sum/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
