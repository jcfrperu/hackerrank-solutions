package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the kangaroo function below.
func kangaroo(x1 float64, v1 float64, x2 float64, v2 float64) string {

	// problem's condition: v2 > v1
	if v2 >= v1 {
		return "NO"
	}

	// position must be integer
	position := (x2 - x1) / (v1 - v2)
	positionInt := float64(int(position))

	// if has decimal part
	if position-positionInt > 0 {
		return "NO"
	}

	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(strings.NewReader("0 2 5 3"), 1024*1024)
	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Stdout, error(nil)
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	x1V1X2V2 := strings.Split(readLine(reader), " ")

	x1Temp, err := strconv.ParseInt(x1V1X2V2[0], 10, 64)
	checkError(err)
	x1 := int(x1Temp)

	v1Temp, err := strconv.ParseInt(x1V1X2V2[1], 10, 64)
	checkError(err)
	v1 := int(v1Temp)

	x2Temp, err := strconv.ParseInt(x1V1X2V2[2], 10, 64)
	checkError(err)
	x2 := int(x2Temp)

	v2Temp, err := strconv.ParseInt(x1V1X2V2[3], 10, 64)
	checkError(err)
	v2 := int(v2Temp)

	result := kangaroo(float64(x1), float64(v1), float64(x2), float64(v2))

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
