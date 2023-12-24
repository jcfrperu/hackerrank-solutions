// *   hackerrank.com DOES NOT SUPPORT TO IMPORT CODE FROM GITHUB (YET).        */
// *   TO SUBMIT YOUR CODE YOU WILL NEED:                                       */
// *   1. include all functions in this template file (copy whole file)         */
// *   2. include your solution() method + methods related to your solution     */
// ******************************************************************************/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*******************************************************************************/
/***************                    SOLUTION                     ***************/
/*******************************************************************************/
func solution(lines []string, writer *bufio.Writer) {
	//TODO: put your code here
	fmt.Fprintf(writer, "%d", 0)
}

func main() {
	RunHackerRank(solution) // config for hackerrank.com
}

/*******************************************************************************/
/***********                  UTILITY METHODS ONLY                   ***********/
/***********     github.com/jcfrperu/go-competitive-programming      ***********/
/*******************************************************************************/

func RunHackerRank(funcToRun func([]string, *bufio.Writer)) {
	run(funcToRun, os.Stdin, os.Getenv("OUTPUT_PATH")) // config for hackerrank.com
}

func run(solutionFunc func([]string, *bufio.Writer), inputReader io.Reader, outFilePath string) {
	getWriter := func(filePath string) (*os.File, *bufio.Writer) {
		if len(filePath) > 0 {
			file, err := os.Create(filePath)
			CheckError(err)
			return file, bufio.NewWriterSize(file, 16*1024*1024)
		}
		return os.Stdout, bufio.NewWriterSize(os.Stdout, 16*1024*1024)
	}

	closeFile := func(file *os.File) {
		err := file.Close()
		CheckError(err)
	}

	flushWriter := func(writer *bufio.Writer) {
		err := writer.Flush()
		CheckError(err)
	}

	getLines := func(input io.Reader) []string {
		reader := bufio.NewReaderSize(input, 16*1024*1024)
		lines := make([]string, 0, 128)
		var str, _, err = reader.ReadLine()
		for err != io.EOF {
			lines = append(lines, strings.TrimRight(string(str), "\r\n"))
			str, _, err = reader.ReadLine()
		}
		return lines
	}

	file, writer := getWriter(outFilePath)
	defer closeFile(file)
	defer flushWriter(writer)
	lines := getLines(inputReader)
	solutionFunc(lines, writer)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func ParseInt(s string) int {
	value, err := strconv.ParseInt(Trim(s), 10, 32)
	CheckError(err)
	return int(value)
}

func ParseLong(s string) int64 {
	value, err := strconv.ParseInt(Trim(s), 10, 64)
	CheckError(err)
	return value
}

func Split(line string, separator string) []string {
	return strings.Split(line, separator)
}

func SplitInts(line string, separator string) []int {
	items := Split(line, separator)
	list := make([]int, 0, len(items))
	for i := range items {
		item := Trim(items[i])
		if item != "" {
			list = append(list, ParseInt(item))
		}
	}
	return list
}

func SplitLongs(line string, separator string) []int64 {
	items := Split(line, separator)
	list := make([]int64, 0, len(items))
	for i := range items {
		item := Trim(items[i])
		if item != "" {
			list = append(list, ParseLong(item))
		}
	}
	return list
}

func SplitGetAt(line string, separator string, index int) string {
	split := Split(line, separator)
	return split[index]
}

func SplitGetIntAt(line string, separator string, index int) int {
	split := Split(line, separator)
	return ParseInt(split[index])
}

func SplitGetLongAt(line string, separator string, index int) int64 {
	split := Split(line, separator)
	return ParseLong(split[index])
}

func Frequencies[T int | int64 | float64 | string](list []T) (map[T]int, []T, int, int) {
	// creating map of frequencies and find indexes for max and min values
	minIndex, maxIndex := 0, 0
	frequencies := make(map[T]int, len(list))
	for i := range list {
		frequencies[list[i]]++
		if list[i] < list[minIndex] {
			minIndex = i
		}
		if list[i] > list[maxIndex] {
			maxIndex = i
		}
	}
	// list of items without repetition
	items := make([]T, 0, len(frequencies))
	for key := range frequencies {
		items = append(items, key)
	}
	// sort list of items by frequency (highest to lowest)
	sort.Slice(items, func(i, j int) bool {
		return frequencies[items[i]] > frequencies[items[j]]
	})
	return frequencies, items, minIndex, maxIndex
}

func Find[T int | int64 | float64 | string](list []T, value T) int {
	for i := range list {
		if list[i] == value {
			return i
		}
	}
	return -1
}

func FindLast[T int | int64 | float64 | string](list []T, value T) int {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == value {
			return i
		}
	}
	return -1
}
