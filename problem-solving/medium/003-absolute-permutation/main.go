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
 * Complete the 'absolutePermutation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 */

func absolutePermutation(n int32, k int32) []int32 {
	// Write your code here
	var (
		arr = make([]int32, n)
		tmp = make(map[int32]int32, n)
		abs = func(v int32) int32 {
			if v < 0 {
				return -v
			}
			return v
		}

		available = func(value int32) int32 {
			av := abs(value)
			if av < 1 || av > n || tmp[av] > 0 {
				return 0
			}

			return av
		}
	)

	for i := int32(0); i < n; i++ {
		arr[i] = i + 1
	}

	if k == 0 {
		return arr
	}

	for _, key := range arr {
		// fmt.Printf("%2d: %d ", key, tmp[key])
		if tmp[key] > 0 {
			// fmt.Printf("continue\n")
			continue
		}

		minusValue := available(key - k)
		plusValue := available(key + k)
		// fmt.Printf("[%2d,%2d] => ", plusValue, minusValue)

		if minusValue == 0 && plusValue == 0 {
			return []int32{-1}
		}

		value := int32(0)

		switch {
		case plusValue > 0 && plusValue != key:
			value = plusValue
		case minusValue > 0 && minusValue != key:
			value = minusValue
		}

		tmp[value] = key
		tmp[key] = value

		// fmt.Printf("%d\n", value)
	}

	ret := make([]int32, 0, n)
	for i := int32(1); i <= n; i++ {
		ret = append(ret, tmp[i])
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		result := absolutePermutation(n, k)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

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
