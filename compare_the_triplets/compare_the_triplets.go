package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {

	logAI32("a", a)
	logAI32("b", b)

	result := []int32{0, 0}

	for i, _ := range a {
		logI("i", i)

		if a[i] > b[i] {
			logS("entro al if", "")
			result[0] += 1
		} else if b[i] > a[i] {
			logS("entro al else", "")
			result[1] += 1
		}
	}

	logAI32("result", result)

	return result
}

var DEBUG bool = false

func logI(prefix string, value int) {
	if DEBUG {
		fmt.Printf("[INFO] "+prefix+"=%v\n", value)
	}
}

func logAI32(prefix string, value []int32) {
	if DEBUG {
		fmt.Printf("[INFO] "+prefix+"=%v\n", value)
	}
}

func logS(prefix string, value string) {
	if DEBUG {
		fmt.Printf("[INFO] "+prefix+"=%v\n", value)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var stdout *os.File
	var err error

	if DEBUG {
		stdout, err = os.Stdout, error(nil)
	} else {
		stdout, err = os.Create(os.Getenv("OUTPUT_PATH"))
	}
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < 3; i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var b []int32

	for i := 0; i < 3; i++ {
		bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	result := compareTriplets(a, b)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
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
