package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"01.gritlab.ax/git/mamberla/guess-it-1/ms"
)

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func guessNextRange(nums []int) string {
	if len(nums) == 0 {
		return "no data"
	}
	numsF := []float64{}
	for _, n := range nums {
		numsF = append(numsF, float64(n))
	}

	rng := [2]int{}

	if len(nums) == 1 {
		rng[0] = int(float64(nums[0]) - ms.Abs(float64(nums[0])))
		rng[1] = int(float64(nums[0]) + ms.Abs(float64(nums[0])))
	} else {
		rng[0] = int(ms.Average(numsF) - ms.Abs(ms.StandardDeviation(numsF)))
		rng[1] = int(ms.Average(numsF) + ms.Abs(ms.StandardDeviation(numsF)))
	}

	return fmt.Sprintf("%v %v", rng[0], rng[1])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "exit" || txt == "quit" {
			break
		}

		if !isInt(txt) {
			fmt.Println("Not an integer:", txt)
		} else {
			num, err := strconv.Atoi(txt)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			numbers = append(numbers, num)
			fmt.Println(guessNextRange(numbers))
		}
	}
}
