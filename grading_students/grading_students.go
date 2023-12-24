package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	grades := make([]int, 0)
	for _, line := range lines[1:] {
		grades = append(grades, ParseInt(line))
	}

	var roundedGrades = make([]int, len(grades))
	for i, grade := range grades {
		roundedGrades[i] = grade

		if grade >= 38 && grade%5 != 0 {
			multipleOfFive := grade + 1
			for multipleOfFive%5 != 0 {
				multipleOfFive++
			}
			if multipleOfFive-grade < 3 {
				roundedGrades[i] = multipleOfFive
			}
		}
	}

	for _, grade := range roundedGrades {
		fmt.Fprintf(writer, "%d\n", grade)
	}
}

func main() {
	//RunWithFile(solution, "grading_students/testcases/001-in.txt")
	RunWithString(solution, "4\n73\n67\n38\n33")
}

// https://www.hackerrank.com/challenges/grading/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
