package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	var (
		re      = regexp.MustCompile(`([01][0-9]):([0-5][0-9]):([0-5][0-9])([AP]M)`)
		matches = re.FindStringSubmatch(s)
	)

	if len(matches) != 5 {
		return fmt.Sprintf("%d", len(matches))
	}

	hours, err := strconv.Atoi(matches[1])
	if err != nil {
		return ""
	}

	if hours == 12 {
		hours = 0
	}

	if matches[4] == "PM" {
		hours += 12
	}

	return fmt.Sprintf("%02d:%s:%s", hours, matches[2], matches[3])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

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
