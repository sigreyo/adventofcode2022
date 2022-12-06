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
			string(stream[i+4]),
			string(stream[i+5]),
			string(stream[i+6]),
			string(stream[i+7]),
			string(stream[i+8]),
			string(stream[i+9]),
			string(stream[i+10]),
			string(stream[i+11]),
			string(stream[i+12]),
			string(stream[i+13]),
		}

		sort.Strings(temp)
		if unique.StringsAreUnique(temp) {
			fmt.Println(i + 14)
			break
		}

	}
}
