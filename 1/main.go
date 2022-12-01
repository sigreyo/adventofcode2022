package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var calsPerPerson []int
	var allCals []int

	for scanner.Scan() {
		tempAdd, _ := strconv.Atoi(scanner.Text())
		allCals = append(allCals, tempAdd)
	}

	var tempAdd int
	for _, cals := range allCals {
		if cals == 0 {
			calsPerPerson = append(calsPerPerson, tempAdd)
			tempAdd = 0
		}
		tempAdd += cals
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calsPerPerson)))

	mostCals := calsPerPerson[0]
	fmt.Printf("most calories: %d\n", mostCals)

	top3 := 0
	for index, cals := range calsPerPerson {
		if index == 3 {
			break
		}

		top3 += cals

	}
	fmt.Printf("top3 has %d calories\n", top3)
}
