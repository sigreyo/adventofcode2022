package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	structure := make(map[string][]string)
	path := ""
	tempPath := ""
	// total := 0
	for _, value := range inputs {
		if value == "$ ls" {
			continue
		}
		structure[path] = append(structure[path], value)

		if strings.Contains(value, "$ cd") && !strings.Contains(value, "..") {
			tempPath = "/" + value[5:]
			path += tempPath

		}
		if strings.Contains(value, "$ cd ..") {
			temp := strings.Split(path, "/")
			path = strings.Join(temp[:len(temp)-1], "/")
		}
		// find := []string{}
		// for _, size := range structure[path] {
		// 	re := regexp.MustCompile("[0-9]+")
		// 	temp := re.FindAllString(size, -1)

		// 	for _, v := range temp {
		// 		check10k, _ := strconv.Atoi(v)
		// 		if check10k < 100000 {
		// 			find = append(find, strconv.Itoa(check10k))
		// 		}
		// 	}
		// }
		// structure[path] = find
		// fmt.Println(structure[path])

	}
	// for key, value := range structure {
	// 	fmt.Println(key, value)
	// }
	// size per folder <100000
	find := []string{}
	for path, value := range structure {
		find = []string{}
		for _, v := range value {
			re := regexp.MustCompile("[0-9]+")
			temp := re.FindAllString(v, -1)
			for _, v := range temp {

				find = append(find, v)

			}
		}
		value = find

		totalSize := 0
		for i := range value {
			temp, _ := strconv.Atoi(value[i])
			totalSize += temp

		}
		// if totalSize > 0 && totalSize < 100000 {
		// 	// fmt.Println(path, value)
		// }
		fmt.Println(path, value)

	}
	// for path, value := range structure {
	// 	fmt.Println(path, value)

	// }

	// totalSize := 0

	// for level, value := range structure {
	// 	re := regexp.MustCompile("[0-9]+")
	// 	values := re.FindAllString(v, -1)

	// }

	// keys := make([]int, len(structure))
	// for key := range structure {
	// 	keys[key] = key
	// 	// fmt.Println(key, value)
	// }
	// sort.Ints(keys)

	// for i := range keys {
	// 	i = len(keys) - 1 - i
	// 	// fmt.Printf("level: %d %s\n", i, structure[i])
	// 	for _, value := range structure[i] {
	// 		if strings.Contains() {

	// 		}
	// 	}

	// }
	// for i := range structure {
	// 	i = len(structure) - 1 - i
	// 	fmt.Println(i)

	// }

}
func addTotal(slice []string) string {
	total := 0
	for _, v := range slice {
		toAdd, _ := strconv.Atoi(v)
		total += toAdd
	}
	return strconv.Itoa(total)
}
