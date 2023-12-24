package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	year := ParseInt(lines[0])

	isJulianCalendar := false
	if year <= 1917 {
		isJulianCalendar = true
	}

	isLeap := false
	if isJulianCalendar {
		isLeap = year%4 == 0
	} else {
		isLeap = (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	}

	daysPerMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year == 1918 {
		daysPerMonth[1] = 28 - 14 + 1 // special case, 1918, February had fewer days since it started the 14th
	} else if isLeap {
		daysPerMonth[1] = 29
	}

	sum := 0
	month := 0
	for sum < 256 {
		sum += daysPerMonth[month]
		month++
	}

	// 256 day is always be in september (9th month)
	var day = 256 - (sum - daysPerMonth[month-1])

	fmt.Fprintf(writer, "%02d.%02d.%d", day, month, year)
}

func main() {
	//RunWithFile(solution, "day_of_the_programmer/testcases/001-in.txt")
	RunWithString(solution, "2016")
}

// https://www.hackerrank.com/challenges/day-of-the-programmer/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
