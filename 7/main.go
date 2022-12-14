package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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
		if scanner.Text() != "$ ls" {
			inputs = append(inputs, scanner.Text())
		}
	}

	structure := make(map[string]int)
	// structureWithInt := make(map[string]int)

	path := ""
	tempPath := ""
	total := 0
	for index, value := range inputs {
		if index == 40 {
			// break
		}
		if value == "$ ls" {
			continue
		}
		re := regexp.MustCompile("[0-9]+")
		temp := re.FindAllString(value, -1)

		//add total file values
		ints := []int{}
		total = 0
		for i := range temp {
			tempInt, _ := strconv.Atoi(temp[i])
			ints = append(ints, tempInt)
			total += ints[i]
		}

		structure[path] += total
		if strings.Contains(value, "$ cd") && !strings.Contains(value, "..") {

			tempPath = "/" + value[5:]
			path += tempPath

		}
		if strings.Contains(value, "$ cd ..") {
			temp := strings.Split(path, "/")
			path = strings.Join(temp[:len(temp)-1], "/")

		}

	}

	keys := make([]string, 0, len(structure))
	for k := range structure {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// skip := ""

	// var slice []string
	// totPerPath := 0
	// var grandtotal []string
	for i, path := range keys {
		if i == 30 {
			// break
		}

		fmt.Println(path, structure[path])
	}
}

// for _, v := range grandtotal {
// 	fmt.Println(v)
// }

func addTotal(slice []string) string {
	total := 0
	for _, v := range slice {
		toAdd, _ := strconv.Atoi(v)
		total += toAdd
	}
	return strconv.Itoa(total)
}
