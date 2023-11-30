// *   hackerrank.com DOES NOT SUPPORT TO IMPORT CODE FROM GITHUB (YET).        */
// *   TO SUBMIT YOUR CODE YOU WILL NEED:                                       */
// *   1. INCLUDE all functions in this template file (copy whole file)         */
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

func ParseInt(s string) int {
	value, err := strconv.ParseInt(s, 10, 32)
	CheckError(err)
	return int(value)
}

func Out(writer *bufio.Writer, data string) {
	PrintOut(writer, data, true, true)
}

func PrintOut(writer *bufio.Writer, data string, trimLeft bool, trimRight bool) {
	dataToPrint := data
	if trimLeft {
		dataToPrint = strings.TrimLeft(dataToPrint, " ")
	}
	if trimRight {
		dataToPrint = strings.TrimRight(dataToPrint, " ")
	}
	_, err := fmt.Fprintf(writer, dataToPrint)
	CheckError(err)
}

func Split(line string, separator string) []string {
	return strings.Split(line, separator)
}

func SplitInts(line string, separator string) []int {
	items := Split(line, separator)
	list := make([]int, 0, len(items))
	for _, item := range items {
		list = append(list, ParseInt(item))
	}
	return list
}

func GetSplitInt(line string, separator string, part int) int {
	split := Split(line, separator)
	return ParseInt(split[part])
}

func FrequenciesInt(list []int) ([]int, map[int]int) {
	// creating map of frequencies
	freqMap := make(map[int]int, len(list))
	for _, item := range list {
		freqMap[item] += 1
	}
	// sorted list of items without repetition
	sortedItems := make([]int, 0, len(freqMap))
	for key, _ := range freqMap {
		sortedItems = append(sortedItems, key)
	}
	sort.Ints(sortedItems)
	return sortedItems, freqMap
}

func ExecInt(operation string, list ...int) int {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}

func ExecIntList(operation string, list []int) int {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}
