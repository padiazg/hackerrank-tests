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
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	// Write your code here
	var (
		isLeapYearJulian = func(year int32) bool {
			return (year % 4) == 0
		}

		isLeapYearGregorian = func(year int32) bool {
			return ((year % 400) == 0) || ((year%4) == 0) && ((year%100) != 0)
		}

		daysTillAugust = 243
	)

	if year == 1918 {
		return "26.08.1918"
	}

	if (year < 1918 && isLeapYearJulian(year)) || (year > 1918 && isLeapYearGregorian(year)) {
		daysTillAugust++
	}

	day := 256 - daysTillAugust

	return fmt.Sprintf("%d.09.%d", day, year)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

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
