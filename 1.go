package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan suara (pisahkan dengan spasi, akhiri dengan 0):")
	scanner.Scan()
	input := scanner.Text()

	votes := strings.Fields(input)
	validVotes := []int{}
	voteCount := make(map[int]int)

	for _, vote := range votes {
		num, err := strconv.Atoi(vote)
		if err != nil {
			continue // Lewati jika tidak bisa diubah menjadi angka
		}
		if num == 0 {
			break // Hentikan jika menemukan 0
		}
		if num >= 1 && num <= 20 {
			validVotes = append(validVotes, num)
			voteCount[num]++
		}
	}

	totalVotes := len(votes) - 1 // Mengurangi 1 untuk suara 0
	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", len(validVotes))

	for candidate, count := range voteCount {
		if count > 0 {
			fmt.Printf("Calon %d: %d suara\n", candidate, count)
		}
	}
}
