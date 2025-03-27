package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	// Write your code here

	// order the list so we can work secuentially
	slices.Sort(a)

	var (
		subarrays = [][]int32{{a[0]}}
		index     int

		check = func(v int32) bool {
			for _, sav := range subarrays[index] {
				diff := v - sav
				if diff < 0 || diff > 1 {
					return false
				}
			}

			return true
		}

		longest int32
	)

	for i := 1; i < len(a); i++ {
		if check(a[i]) {
			subarrays[index] = append(subarrays[index], a[i])
		} else {
			subarrays = append(subarrays, []int32{})
			index = len(subarrays) - 1
			subarrays[index] = []int32{a[i]}
		}
	}

	for _, sa := range subarrays {
		len := int32(len(sa))
		if len > longest {
			longest = len
		}
	}

	return longest
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

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
