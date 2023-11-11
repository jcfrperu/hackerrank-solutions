package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
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

func main() {
	//reader := bufio.NewReaderSize(strings.NewReader("12:05:39AM"), 1024 * 1024)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//outputFile, err := os.Stdout, error(nil)
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
