package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

func solution(lines []string, writer *bufio.Writer) {
	s := lines[0]

	split := strings.Split(s, ":")

	hours := ParseInt(split[0])
	minutes := ParseInt(split[1])
	seconds := ParseInt(strings.ReplaceAll(strings.ReplaceAll(split[2], "AM", ""), "PM", ""))

	if strings.Contains(split[2], "PM") && hours < 12 {
		hours = hours + 12
	} else if strings.Contains(split[2], "AM") && hours >= 12 {
		hours = hours - 12
	}

	fmt.Fprintf(writer, "%02d:%02d:%02d", hours, minutes, seconds)
}

func main() {
	RunWithFile(solution, "time_conversion/testcases/001-in.txt")
	//RunWithString(solution, "07:05:45PM")
}

// https://www.hackerrank.com/challenges/time-conversion/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
