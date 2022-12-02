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
		totalScore += riggedScore(line[0], line[1])
		// fmt.Printf("o:%s, y:%s\n", line[0], line[1])
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

func riggedScore(opp, you string) int {

	outcome := 0
	switch opp {
	case oppRock:
		switch you {
		case "X":
			outcome = 0 + shapeScore(youScissors)
		case "Y":
			outcome = 3 + shapeScore(youRock)
		case "Z":
			outcome = 6 + shapeScore(youPaper)
		}
	case oppPaper:
		switch you {
		case "X":
			outcome = 0 + shapeScore(youRock)
		case "Y":
			outcome = 3 + shapeScore(youPaper)
		case "Z":
			outcome = 6 + shapeScore(youScissors)
		}
	case oppScissors:
		switch you {
		case "X":
			outcome = 0 + shapeScore(youPaper)
		case "Y":
			outcome = 3 + shapeScore(youScissors)
		case "Z":
			outcome = 6 + shapeScore(youRock)
		}
	}

	return outcome
}

func shapeScore(shape string) int {
	switch shape {
	case oppRock, youRock:
		return 1
	case oppPaper, youPaper:
		return 2
	case oppScissors, youScissors:
		return 3
	}
	return 0
}
