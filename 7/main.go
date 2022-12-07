package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	inputs := []string{}
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	structure := make(map[int][]string)
	levelTracker := 0
	for _, value := range inputs {
		structure[levelTracker] = append(structure[levelTracker], "-"+value)

		if strings.Contains(value, "$ cd") && !strings.Contains(value, "..") {
			levelTracker++
		}
		if strings.Contains(value, "$ cd ..") {
			levelTracker--
		}

	}

	// totalSize := 0

	// for level, value := range structure {
	// 	re := regexp.MustCompile("[0-9]+")
	// 	values := re.FindAllString(v, -1)

	// }

	keys := make([]int, len(structure))
	for key := range structure {
		keys[key] = key
		// fmt.Println(key, value)
	}
	sort.Ints(keys)

	for i := range keys {
		i = len(keys) - 1 - i
		// fmt.Printf("level: %d %s\n", i, structure[i])
		for _, value := range structure[i] {
			if strings.Contains() {

			}
		}

	}
	// for i := range structure {
	// 	i = len(structure) - 1 - i
	// 	fmt.Println(i)

	// }

}
