package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	n := ParseInt(lines[0])

	for iter := 1; iter <= n; iter++ {
		A := SplitInts(lines[iter*2], " ")

		i := 0
		for canBeIterated(A) && !isSorted(A) {
			if mustBeRotate(A, i) {
				rotate(A, i)
			} else {
				i++
				if i > len(A)-3 {
					i = 0
				}
			}
		}

		if isSorted(A) {
			fmt.Fprintf(writer, "YES\n")
		} else {
			fmt.Fprintf(writer, "NO\n")
		}
	}
}

func rotate(a []int, i int) {
	if i >= 0 && i <= len(a)-3 {
		temp := a[i]
		a[i] = a[i+1]
		a[i+1] = a[i+2]
		a[i+2] = temp
	}
}

func mustBeRotate(a []int, i int) bool {
	if i >= 0 && i <= len(a)-3 {
		return a[i] > a[i+1] || a[i] > a[i+2]
	}
	return false
}

func isSorted(a []int) bool {
	// sorted from lowest to highest
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func canBeIterated(a []int) bool {
	for i := 0; i <= len(a)-3; i++ {
		if mustBeRotate(a, i) {
			return true
		}
	}
	return false
}

func main() {
	//RunWithFile(solution, "larrys_array/testcases/001-in.txt")
	RunWithString(solution, "3\n3\n3 1 2\n4\n1 3 4 2\n5\n1 2 3 5 4")
}

// https://www.hackerrank.com/challenges/larrys-array/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
