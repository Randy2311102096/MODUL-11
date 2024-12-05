package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	votes := strings.Fields(input)
	validVotes := []int{}
	voteCount := make(map[int]int)

	for _, vote := range votes {
		num, err := strconv.Atoi(vote)
		if err == nil && num >= 1 && num <= 20 {
			validVotes = append(validVotes, num)
			voteCount[num]++
		}
		if num == 0 {
			break
		}
	}

	fmt.Printf("Suara masuk: %d\n", len(votes)-1)
	fmt.Printf("Suara sah: %d\n", len(validVotes))

	// Menentukan pemenang
	type candidate struct {
		number int
		votes  int
	}
	var candidates []candidate

	for num, count := range voteCount {
		candidates = append(candidates, candidate{number: num, votes: count})
	}

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].votes == candidates[j].votes {
			return candidates[i].number < candidates[j].number
		}
		return candidates[i].votes > candidates[j].votes
	})

	if len(candidates) > 0 {
		fmt.Printf("Ketua RT: %d\n", candidates[0].number)
	}
	if len(candidates) > 1 {
		fmt.Printf("Wakil ketua: %d\n", candidates[1].number)
	}
}