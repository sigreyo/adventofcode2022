package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/mpvl/unique"
)

func main() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	stream := ""
	for scanner.Scan() {
		stream += scanner.Text()
	}

	for i := range stream {
		temp := []string{string(stream[i]),
			string(stream[i+1]),
			string(stream[i+2]),
			string(stream[i+3]),
		}

		sort.Strings(temp)
		if unique.StringsAreUnique(temp) {
			fmt.Println(i + 4)
			break
		}

	}
}
