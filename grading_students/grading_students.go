package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func gradingStudents(grades []int) []int {

	var gradesRounded = make([]int, len(grades), len(grades))

	for i, grade := range grades {

		gradesRounded[i] = grade

		if grade >= 38 && grade%5 != 0 {

			multipleOfFive := grade + 1
			for multipleOfFive%5 != 0 {
				multipleOfFive++
			}
			if multipleOfFive-grade < 3 {
				gradesRounded[i] = multipleOfFive
			}
		}
	}

	return gradesRounded
}

func solution(lines []string, writer *bufio.Writer) {

	Out(writer, fmt.Sprintf("%d", 4))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
