package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strconv"
	"strings"
)

func timeConversion(s string) string {

	split := strings.Split(s, ":")

	horas, _ := strconv.Atoi(split[0])
	minutos, _ := strconv.Atoi(split[1])
	segundos, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(split[2], "AM", ""), "PM", ""))

	if strings.Contains(split[2], "PM") && horas < 12 {
		horas = horas + 12
	} else if strings.Contains(split[2], "AM") && horas >= 12 {
		horas = horas - 12
	}

	//result := strconv.Itoa(horas) + ":" + strconv.Itoa(minutos) + ":" + strconv.Itoa(segundos)
	result := fmt.Sprintf("%02d:%02d:%02d", horas, minutos, segundos)

	return result
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
