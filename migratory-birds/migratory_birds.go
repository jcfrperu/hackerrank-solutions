package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int) int {

	var acum = []int{0, 0, 0, 0, 0, 0}

	for i := range arr {
		acum[arr[i]]++
	}

	var max = arr[0]
	for i := 1; i < len(acum); i++ {
		if acum[i] > max {
			max = acum[i]
		}
	}

	var number = 0
	for i := 1; i < len(acum); i++ {
		if acum[i] == max {
			number = i
			break
		}
	}

	return number
}

func main() {
	//reader := bufio.NewReaderSize(strings.NewReader("6\n1 4 4 4 5 3"), 16*1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Stdout, error(nil)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
