package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

// final version, won't trigger a timeout and pass the test
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	var (
		history      = make([]int32, 0, len(player))
		rankedUnique = func() []int32 {
			var (
				ranks    = ranked[:1]
				previous = ranks[0]
			)

			for _, score := range ranked {
				if score != previous {
					ranks = append(ranks, score)
				}
				previous = score
			}

			return ranks
		}()
	)

	// fmt.Printf("%d => %d\n", len(ranked), len(rankedUnique))

	for _, score := range player {
		rank := sort.Search(len(rankedUnique), func(i int) bool { return rankedUnique[i] <= score })
		history = append(history, int32(rank+1))
	}

	return history
}

// triggers timeout, wont pass two tests
func climbingLeaderboard_c(ranked []int32, player []int32) []int32 {
	// Write your code here
	history := []int32{}

	unique := func() []int32 {
		ranks := ranked[:1]
		previous := ranks[0]
		for _, score := range ranked {
			if score != previous {
				ranks = append(ranks, score)
			}
			previous = score
		}
		return ranks
	}

	_ranked := unique()
	fmt.Printf("%d => %d\n", len(ranked), len(_ranked))

	// for _, score := range player {
	for z := 0; z < len(player); z++ {
		currentRank := int32(0)
		currentScore := _ranked[0] + 1

		for i := 0; i < len(_ranked); i++ {
			if player[z] >= _ranked[i] {
				break
			}
			if _ranked[i] < currentScore {
				currentRank++
				currentScore = _ranked[i]
			}
		}

		// return currentRank + 1
		history = append(history, currentRank+1)
	}

	return history
}

// triggers timeout, wont pass two tests
func climbingLeaderboard_b(ranked []int32, player []int32) []int32 {
	// Write your code here
	var (
		history = []int32{}

		rank = func(score int32) int32 {
			currentRank := int32(0)
			currentScore := ranked[0] + 1

			for i := 0; i < len(ranked); i++ {
				s := ranked[i]
				if s < currentScore {
					currentRank++
					currentScore = s
				}
				if score >= s {
					return currentRank
				}
			}

			return currentRank + 1
		}
	)

	for _, score := range player {
		history = append(history, rank(score))
	}

	return history
}

// triggers timeout, wont pass two tests
func climbingLeaderboard_a(ranked []int32, player []int32) []int32 {
	// Write your code here
	var (
		history = []int32{}

		rank = func(score int32) int32 {
			var (
				currentRank = int32(0)
				tmp_ranked  = make([]int32, len(ranked))
			)

			// make a deep copy of the original ranked
			tmp_ranked = append(tmp_ranked, ranked...)

			// add
			tmp_ranked = append(tmp_ranked, score)
			slices.Sort(tmp_ranked)
			currentScore := tmp_ranked[len(tmp_ranked)-1] + 1

			for i := len(tmp_ranked); i > 0; i-- {
				s := tmp_ranked[i-1]
				if s < currentScore {
					currentRank++
					currentScore = s
				}

				if s == score {
					return currentRank
				}

			}

			return 0
		}
	)

	for _, score := range player {
		history = append(history, rank(score))
	}

	return history
}

func climbingLeaderboard2(ranked []int32, player []int32) []int32 {
	ranks := ranked[:1]
	last := ranks[0]
	for _, score := range ranked[1:] {
		if score != last {
			ranks = append(ranks, score)
		}
		last = score
	}

	climb := make([]int32, 0, len(player))
	for _, score := range player {
		rank := sort.Search(
			len(ranks),
			func(i int) bool { return ranks[i] <= score },
		)
		climb = append(climb, int32(rank+1))
	}
	return climb
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
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
