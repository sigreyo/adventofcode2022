package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	oppRock     = "A"
	youRock     = "X"
	oppPaper    = "B"
	youPaper    = "Y"
	oppScissors = "C"
	youScissors = "Z"
)

func main() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var totalScore int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		totalScore += determineScore(line[0], line[1])
	}
	fmt.Printf("Total score is: %v\n", totalScore)
}

func determineScore(opp, you string) int {

	outcome := 0
	switch you {
	case youRock:
		switch opp {
		case oppRock:
			outcome = 3
		case oppPaper:
			outcome = 0
		case oppScissors:
			outcome = 6
		}
	case youPaper:
		switch opp {
		case oppRock:
			outcome = 6
		case oppPaper:
			outcome = 3
		case oppScissors:
			outcome = 0
		}
	case youScissors:
		switch opp {
		case oppRock:
			outcome = 0
		case oppPaper:
			outcome = 6
		case oppScissors:
			outcome = 3
		}
	}
	return outcome + shapeScore(you)
}

func shapeScore(shape string) int {
	switch shape {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return 0
}
