package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
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
	for key := range boxSet {
		reverse(boxSet[key])
	}
	counter = 0
	actions := readActions()

	for _, v := range actions {
		counter++
		re := regexp.MustCompile("[0-9]+")
		instructions := re.FindAllString(v, -1)
		boxSet = removeBlanks(boxSet)

		amountToMove, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		to, _ := strconv.Atoi(instructions[2])

		temp := boxSet[from][len(boxSet[from])-amountToMove:]
		boxSet[from] = slices.Delete(boxSet[from], len(boxSet[from])-amountToMove, len(boxSet[from]))
		boxSet[to] = append(boxSet[to], temp...)

		// for i := range temp {
		// 	i = len(temp) - 1 - i
		// 	boxSet[to] = append(boxSet[to], temp[i])
		// }

	}
	for key, value := range boxSet {
		fmt.Println(key, value[len(value)-1])
	}

}

func removeBlanks(boxSet map[int][]string) map[int][]string {
	for key, value := range boxSet {

		for i, v := range value {
			if v == " " {
				boxSet[key] = value[0:i:i]
				break
			}
		}

	}
	return boxSet

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
