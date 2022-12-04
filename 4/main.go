package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var counter int
	for scanner.Scan() {
		commaRemoved := strings.Split(scanner.Text(), ",")

		firstElf := fillSlice(strings.Split(commaRemoved[0], "-"))
		secondElf := fillSlice(strings.Split(commaRemoved[1], "-"))

		if checkForOverlap(firstElf, secondElf) {
			counter++
		}
	}
	fmt.Println(counter)
}

func fillSlice(cleaningID []string) []int {
	var temp []int

	//convert to int slice
	for i := range cleaningID {
		strToInt, _ := strconv.Atoi(cleaningID[i])
		temp = append(temp, strToInt)
	}

	//create new slice with all values
	var fullSlice []int
	for i := temp[0]; i <= temp[1]; i++ {

		fullSlice = append(fullSlice, i)
	}

	return fullSlice
}

func checkForOverlap(first, second []int) bool {

	// if first[0] <= second[0] && first[len(first)-1] >= second[len(second)-1] {
	// 	return true
	// }
	// if second[0] <= first[0] && second[len(second)-1] >= first[len(first)-1] {
	// 	return true
	// }
	for _, v := range first {
		if slices.Contains(second, v) {
			return true
		}
	}
	return false

}
