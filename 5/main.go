package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("crateorigin.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 1
	boxSet := make(map[int][]string)
	for scanner.Scan() {
		if counter == 9 {
			break

		}
		line := strings.Split(scanner.Text(), "")
		line = trimSlice(line)
		for i, v := range line {
			boxSet[i+1] = append(boxSet[i+1], v)
		}
		counter++
	}
	for key, value := range boxSet {
		reverse(boxSet[key])

		for i, v := range value {
			if v == " " {
				value = value[0:i:i]
				break
			}
		}
		// fmt.Println(key, value)

	}
	actions := readActions()

	for _, v := range actions {
		re := regexp.MustCompile("[0-9]+")
		instructions := re.FindAllString(v, -1)

		amountToMove, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		// to, _ := strconv.Atoi(instructions[2])

		tempMove := boxSet[from][len(boxSet[from])-amountToMove : amountToMove]
		fmt.Println(tempMove)
		// boxSet[to] = append(boxSet[from])

		// fmt.Println(amountToMove, from, to)
		// boxSet = move(amountToMove, from, to, boxSet)
	}
	// for key, value := range boxSet {
	// 	fmt.Println(key, value)
	// }

}

func move(amountToMove, from, to int, boxSet map[int][]string) (newSet map[int][]string) {
	var tempMove []string

	for i := range boxSet[from] {
		i = len(boxSet[from]) - i - 1
		if len(tempMove) != amountToMove {
			tempMove = append(tempMove, boxSet[from][i])
		}
	}
	for i := range tempMove {
		boxSet[to] = append(boxSet[to], tempMove[i])
	}
	fmt.Println(tempMove)

	return
}
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func readActions() (res []string) {
	f, err := os.Open("actions.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return
}

func trimSlice(line []string) []string {
	for i := range line {
		if line[i] == "[" || line[i] == "]" {
			line[i] = ""
		}
	}
	for i := range line {
		if i%2 == 0 {
			line[i] = ""
		}
	}
	for i := range line {
		if i%2 == 0 {
			line[i] = ""
		}
	}
	for i := range line {
		if i%2 == 0 {
			line[i] = ""
		}
	}
	var new []string
	for i := 1; i < len(line)-1; i += 4 {
		new = append(new, line[i])

	}

	return new
}
