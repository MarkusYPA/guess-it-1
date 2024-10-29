package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatalln("Give two arguments")
	}

	address := fmt.Sprintf("data_sets/%s/%s.txt", os.Args[1], os.Args[2])

	numsFile, err := os.ReadFile(address)
	if err != nil {
		log.Fatalln(err.Error())
	}
	numsStrings := strings.Split(string(numsFile), "\n")

	ints := []int{}
	for _, ns := range numsStrings {
		n, e := strconv.Atoi(ns)
		if e != nil {
			continue
		}
		ints = append(ints, n)
	}

	for _, v := range ints {
		fmt.Println(v)
	}
}
