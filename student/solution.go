package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"01.gritlab.ax/git/mamberla/guess-it-1/ms"
)

// isInt checks if the provided string is an integer
func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, r := range s {
		if i == 0 && !unicode.IsNumber(r) && r != '-' && r != '+' {
			return false
		}

		if i > 0 && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

// avgAndSD returns the mean of the data plus and minus the standard deviation
func avgAndSD(nums []float64) [2]int {
	rng := [2]int{}

	if len(nums) == 1 {
		rng[0] = ms.RoundToInt(nums[0]) - ms.RoundToInt((ms.Abs(nums[0]))/2.0)
		rng[1] = ms.RoundToInt(nums[0]) + ms.RoundToInt((ms.Abs(nums[0]))/2.0)
	} else {
		rng[0] = ms.RoundToInt(ms.Average(nums) - ms.Abs(ms.StandardDeviation(nums)))
		rng[1] = ms.RoundToInt(ms.Average(nums) + ms.Abs(ms.StandardDeviation(nums)))
	}

	return rng
}

// box returns 1st and 3rd quartile cutting points, often seen in box plots
func box(nums []float64) [2]int {
	rng := [2]int{}

	if len(nums) == 1 {
		rng[0] = ms.RoundToInt(nums[0]) - ms.RoundToInt((ms.Abs(nums[0]))/2.0)
		rng[1] = ms.RoundToInt(nums[0]) + ms.RoundToInt((ms.Abs(nums[0]))/2.0)
	} else {
		rng[0] = ms.RoundToInt(ms.Quarters(nums)[1])
		rng[1] = ms.RoundToInt(ms.Quarters(nums)[3])
	}

	return rng
}

// guessNextRange makes a guess about in which range the next number will be
func guessNextRange(nums []int) string {

	if len(nums) == 0 {
		return "no data"
	}

	numsF := []float64{}
	for _, n := range nums {
		numsF = append(numsF, float64(n))
	}

	//rng := avgAndSD(numsF)
	rng := box(numsF)

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
