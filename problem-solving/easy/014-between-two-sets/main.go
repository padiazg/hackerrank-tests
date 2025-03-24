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
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var (
		factors     = map[int32]int32{}
		count       = int32(0)
		findFactors = func(factors map[int32]int32, values []int32) map[int32]int32 {
			for _, value := range values {
				for a := int32(1); a <= value; a++ {
					if value%a == 0 {
						factors[a] = a
					}
				}
			}

			return factors
		}
	)

	factors = findFactors(factors, a)
	// fmt.Printf("a = %v\n    factors: %+v\n", a, factors)
	factors = findFactors(factors, b)
	// fmt.Printf("b = %v\n    factors: %+v\n", b, factors)

	for factor := range factors {
		match_a := true
		match_b := true

		for _, va := range a {
			if factor%va != 0 {
				match_a = false
				break
			}
		}

		for _, vb := range b {
			if vb%factor != 0 {
				match_b = false
				break
			}
		}

		if match_a && match_b {
			// fmt.Printf("match factor=%d\n", factor)
			count++
		}
	}

	// fmt.Printf("count=%d", count)
	return count
}

// func getTotalX(a []int32, b []int32) int32 {
// 	// Write your code here
// 	factors := map[int32]int32{}
// 	matches := map[int32]int32{}

// 	for _, v := range a {
// 		factors[v] = v
// 		factors[v+v] = v + v
// 		for _, x := range a[1:] {
// 			factors[x+v] = x + v
// 		}
// 	}

// 	fmt.Printf("a = %v\n", a)
// 	fmt.Printf("b = %v\n", b)
// 	fmt.Printf("factors = %v\n", factors)

// 	for factor := range factors {
// 		match := true
// 		for _, v := range b {
// 			m := v % factor
// 			fmt.Printf("%d mod %d = %d\n", v, factor, m)
// 			if m != 0 {
// 				match = false
// 				break
// 			}
// 		}
// 		if match {
// 			matches[factor] = factor
// 		}
// 	}

// 	fmt.Printf("matches = %+v count = %d", matches, len(matches))
// 	return int32(len(matches))
// }

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

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
