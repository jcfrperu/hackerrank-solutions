package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func main() {
	reader := bufio.NewReaderSize(strings.NewReader("4\n73\n67\n38\n33"), 16*1024*1024)
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Stdout, error(nil)
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var grades []int

	for i := 0; i < int(gradesCount); i++ {
		gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		gradesItem := int(gradesItemTemp)
		grades = append(grades, gradesItem)
	}

	result := gradingStudents(grades)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
