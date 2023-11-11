package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int, t int, a int, b int, apples []int, oranges []int) {

	var numberApples = 0
	var numberOranges = 0

	for _, da := range apples {
		if a+da >= s && a+da <= t {
			numberApples++
		}
	}

	for _, db := range oranges {
		if b+db <= t && b+db >= s {
			numberOranges++
		}
	}

	fmt.Printf("%d\n%d", numberApples, numberOranges)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	st := strings.Split(readLine(reader), " ")

	sTemp, err := strconv.ParseInt(st[0], 10, 64)
	checkError(err)
	s := int(sTemp)

	tTemp, err := strconv.ParseInt(st[1], 10, 64)
	checkError(err)
	t := int(tTemp)

	ab := strings.Split(readLine(reader), " ")

	aTemp, err := strconv.ParseInt(ab[0], 10, 64)
	checkError(err)
	a := int(aTemp)

	bTemp, err := strconv.ParseInt(ab[1], 10, 64)
	checkError(err)
	b := int(bTemp)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int(nTemp)

	applesTemp := strings.Split(readLine(reader), " ")

	var apples []int

	for i := 0; i < int(m); i++ {
		applesItemTemp, err := strconv.ParseInt(applesTemp[i], 10, 64)
		checkError(err)
		applesItem := int(applesItemTemp)
		apples = append(apples, applesItem)
	}

	orangesTemp := strings.Split(readLine(reader), " ")

	var oranges []int

	for i := 0; i < int(n); i++ {
		orangesItemTemp, err := strconv.ParseInt(orangesTemp[i], 10, 64)
		checkError(err)
		orangesItem := int(orangesItemTemp)
		oranges = append(oranges, orangesItem)
	}

	countApplesAndOranges(s, t, a, b, apples, oranges)
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
