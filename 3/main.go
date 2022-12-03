package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalScore := 0

	for scanner.Scan() {

		strLen := len(scanner.Text())

		firstHalf := strings.Split(scanner.Text()[:strLen/2], "")
		// secondHalf := strings.Split(scanner.Text()[strLen/2:], "")
		secondHalf := scanner.Text()[strLen/2:]

		for index := range firstHalf {
			if strings.Contains(secondHalf, firstHalf[index]) {
				totalScore += prioScore(firstHalf[index])
				break
			}
		}

	}

	fmt.Println(totalScore)
}

func prioScore(letter string) int {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if strings.Contains(lower, letter) {
		return strings.Index(lower, letter) + 1
	} else {
		return strings.Index(upper, letter) + 27
	}

}
