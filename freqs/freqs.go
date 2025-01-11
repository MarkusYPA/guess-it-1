package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	numsMap := make(map[string]int)
	numsSlice := []string{}

	for i := 1; i < 4; i++ {
		for j := 1; j < 6; j++ {
			address := fmt.Sprintf("data_sets/%v/%v.txt", i, j)
			bytes, err := os.ReadFile(address)
			check(err)
			numsSlice = append(numsSlice, strings.Split(string(bytes), "\n")...)
		}
	}

	for _, n := range numsSlice {
		numsMap[n]++
	}

	numsAgain := make([][2]int, 101)
	for key, value := range numsMap {

		if key == "" {
			continue
		}

		keyInt, err := strconv.Atoi(key)
		check(err)

		if keyInt > 200 || keyInt < 100 {
			continue
		}

		numsAgain[keyInt-100] = [2]int{keyInt, value}
	}

	// Sort by frequency if you like
	for i := 0; i < len(numsAgain)-1; i++ {
		for j := i + 1; j < len(numsAgain); j++ {
			if numsAgain[i][1] < numsAgain[j][1] {
				numsAgain[i], numsAgain[j] = numsAgain[j], numsAgain[i]
			}
		}
	}

	for _, numFreq := range numsAgain {
		fmt.Println(numFreq[0], numFreq[1])
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e.Error())
	}
}
