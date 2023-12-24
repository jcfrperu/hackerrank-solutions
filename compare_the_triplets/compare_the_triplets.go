package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	a := SplitInts(lines[0], " ")
	b := SplitInts(lines[1], " ")

	result := []int32{0, 0}
	for i := range a {
		if a[i] > b[i] {
			result[0] += 1
		} else if b[i] > a[i] {
			result[1] += 1
		}
	}

	fmt.Fprintf(writer, "%d %d", result[0], result[1])
}

func main() {
	RunWithFile(solution, "compare_the_triplets/testcases/001-in.txt")
	//RunWithString(solution, "5 6 7\n3 6 10")
}

// https://www.hackerrank.com/challenges/compare-the-triplets/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
