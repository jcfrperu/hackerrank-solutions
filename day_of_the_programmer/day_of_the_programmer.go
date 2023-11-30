package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int) string {

	var isJulianCalendar = false
	if year <= 1917 {
		isJulianCalendar = true
	}

	var isLeap = false
	if isJulianCalendar {
		isLeap = year%4 == 0
	} else {
		isLeap = (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	}

	var daysPerMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year == 1918 {
		daysPerMonth[1] = 28 - 14 + 1 // caso especial en el anio 1918 que febrero tuvo menos dias porque empezo el nro 14
	} else if isLeap {
		daysPerMonth[1] = 29
	}

	var sum = 0
	var month = 0
	for sum < 256 {
		sum += daysPerMonth[month]
		month++
	}

	// el dia 256 siempre sera en setiembre 9th mes
	var day = 256 - (sum - daysPerMonth[month-1])

	return fmt.Sprintf("%02d.%02d.%d", day, month, year)
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
