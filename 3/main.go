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

	// for scanner.Scan() {

	// 	strLen := len(scanner.Text())

	// 	firstHalf := strings.Split(scanner.Text()[:strLen/2], "")
	// 	// secondHalf := strings.Split(scanner.Text()[strLen/2:], "")
	// 	secondHalf := scanner.Text()[strLen/2:]

	// 	for index := range firstHalf {
	// 		if strings.Contains(secondHalf, firstHalf[index]) {
	// 			totalScore += prioScore(firstHalf[index])
	// 			break
	// 		}
	// 	}

	// }
	counter := 1
	var temp []string
	for scanner.Scan() {

		if counter%3 == 0 {
			badge := strings.Split(scanner.Text(), "")

			for index := range badge {
				if strings.Contains(temp[0], badge[index]) && strings.Contains(temp[1], badge[index]) {
					totalScore += prioScore(badge[index])
					temp = nil
					break
				}
			}
		} else {
			temp = append(temp, scanner.Text())
		}

		counter++

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
