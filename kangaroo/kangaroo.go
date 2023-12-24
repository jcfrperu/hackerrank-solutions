package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	x1 := float64(SplitGetIntAt(lines[0], " ", 0))
	v1 := float64(SplitGetIntAt(lines[0], " ", 1))
	x2 := float64(SplitGetIntAt(lines[0], " ", 2))
	v2 := float64(SplitGetIntAt(lines[0], " ", 3))

	// problem's condition: v2 > v1
	if v2 >= v1 {
		fmt.Fprintf(writer, "NO")
		return
	}

	// position must be integer
	var position = (x2 - x1) / (v1 - v2)
	positionInt := float64(int(position))

	// if it has decimal part
	if position-positionInt > 0 {
		fmt.Fprintf(writer, "NO")
		return
	}

	fmt.Fprintf(writer, "YES")
}

func main() {
	//RunWithFile(solution, "kangaroo/testcases/001-in.txt")
	RunWithString(solution, "0 2 5 3")
}

// https://www.hackerrank.com/challenges/kangaroo/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
