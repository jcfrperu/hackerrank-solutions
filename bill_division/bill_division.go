package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

/* ****************************************************************** */
/* **************************   SOLUTION   ************************** */
/* ****************************************************************** */
func solution(lines []string, writer *bufio.Writer) {
	k := SplitIntsGetAt(lines[0], " ", 1)
	bill := SplitInts(lines[1], " ")
	bCharged := ParseInt(lines[2])

	sum := 0
	for i, item := range bill {
		if i == k {
			continue
		}
		sum += item
	}

	bActual := sum / 2
	result := bCharged - bActual

	if result == 0 {
		Out(writer, fmt.Sprintf("Bon Appetit"))
	} else {
		Out(writer, fmt.Sprintf("%d", result))
	}
}

func main() {
	//RunWithFile(solution, "bill_division/testcases/001.input")
	RunWithString(solution, "7 11\n5 15\n3 2\n-2 2 1\n5 -6")
}

// https://www.hackerrank.com/challenges/apple-and-orange/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
