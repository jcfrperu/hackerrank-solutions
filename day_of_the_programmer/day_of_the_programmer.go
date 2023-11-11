package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func main() {
	//reader := bufio.NewReaderSize(strings.NewReader("1918"), 16*1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Stdout, error(nil)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int(yearTemp)

	result := dayOfProgrammer(year)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
