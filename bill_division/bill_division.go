package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	k := SplitGetIntAt(lines[0], " ", 1)
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
		fmt.Fprintf(writer, "Bon Appetit")
	} else {
		fmt.Fprintf(writer, "%d", result)
	}
}

func main() {
	//RunWithFile(solution, "bill_division/testcases/001-in.txt")
	RunWithString(solution, "4 1\n3 10 2 9\n12")
}

// https://www.hackerrank.com/challenges/bon-appetit/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
