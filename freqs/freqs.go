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

	numsAgain := make([]int, 101)
	for key, value := range numsMap {

		if key == "" {
			continue
		}

		keyInt, err := strconv.Atoi(key)
		check(err)

		if keyInt > 200 || keyInt < 100 {
			continue
		}

		numsAgain[keyInt-100] = value
	}

	//sort.Ints(numsAgain)

	for i, n := range numsAgain {
		fmt.Println(i+100, n)
	}
	//fmt.Println(numsAgain)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e.Error())
	}
}
