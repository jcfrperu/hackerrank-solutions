package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	t := ParseInt(lines[0])
	for i := 1; i <= t; i++ {
		n := SplitGetIntAt(lines[i], " ", 0)
		c := SplitGetIntAt(lines[i], " ", 1)
		m := SplitGetIntAt(lines[i], " ", 2)

		chocolatesNumber := calculateChocolatesNumber(n, c, m)
		fmt.Fprintf(writer, "%d\n", chocolatesNumber)
	}
}

func calculateChocolatesNumber(totalMoney int, unitCost int, wrappersGetNewOne int) int {
	chocolates := totalMoney / unitCost

	wrappers := chocolates
	for wrappers >= wrappersGetNewOne {
		division := wrappers / wrappersGetNewOne
		reminder := wrappers % wrappersGetNewOne
		wrappers = division + reminder

		chocolates += division
	}

	return chocolates
}

func main() {
	//RunWithFile(solution, "chocolate_feast/testcases/001-in.txt")
	RunWithString(solution, "3\n10 2 5\n12 4 4\n6 2 2")
	//RunWithString(solution, "1\n15 3 2")
}

// https://www.hackerrank.com/challenges/chocolate-feast/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
