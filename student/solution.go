package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"

	"01.gritlab.ax/git/mamberla/guess-it-1/guessing"
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

// guessNextRange makes a guess about in which range the next number will be
func guessNextRange(nums []int, mult float64) (int, int) {

	if len(nums) == 0 {
		return 0, 0
	}

	numsF := toFloats(nums)

	//rng := guessing.Box(numsF)
	//rng := guessing.AvgAndSD(numsF, mult) // 3* 3/5 + 1* 4/5 with mult 0.67
	//rng := [2]int{nums[len(nums)-1], nums[len(nums)-1]}
	//rng := guessing.SillyGuess(nums)
	//rng := guessing.HardRange()	// 14 / 15 with [2]int{112, 113}
	//rng := guessing.MedToMean(numsF, mult) // 1* 2/5, 2* 3/5, 1* 4/5 with mult 1.0

	rng := guessing.MedAndSD(numsF, mult) // 2* 4/5 with mult 0.7 and 23 data points

	return rng[0], rng[1]
}

func main() {
	pointsOn := flag.Bool("points", false, "To display accrued points at the end, set to true")
	flag.Parse()

	mult := 0.71 //0.67
	if len(flag.Args()) == 1 {
		val, e := strconv.ParseFloat(flag.Arg(0), 64)
		if e != nil {
			log.Fatalln(e.Error())
		}
		mult = val
	}

	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}
	seen := 0
	guess := "100 200"
	r1, r2 := 0, 0
	points := 0

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "exit" || txt == "quit" || txt == "" {
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

			points += getPoints(num, r1, r2)

			// Use old guess and don't use the data point for calculations if it's an outlier
			if seen > 3 && isOutlier(num, numbers) {
				fmt.Println(guess)
				continue
			}

			// Only ever deal with data sets maximum 23 long
			if seen < 23 {
				seen++
				numbers = append(numbers, num)
			} else {
				numbers = append(numbers[1:], num)
			}

			r1, r2 = guessNextRange(numbers, mult)
			guess = fmt.Sprintf("%v %v", r1, r2)

			fmt.Printf("%s\n", guess)
		}
	}
	if *pointsOn {
		fmt.Printf("\n%v\n", points)
	}
}

// getPoints counts points for testing purposes
func getPoints(num, r1, r2 int) int {
	points := 0.0
	rangeWidth := r2 - r1 + 1
	if num >= r1 && num <= r2 {
		points += 800 / float64(rangeWidth)
	}
	return int(math.Round(points))
}

// isOutlier tells if a data point is way off most of the others
func isOutlier(n int, nums []int) bool {
	floats := toFloats(nums)
	return ms.Abs(float64(n)-ms.Average(floats)) > ms.Variance(floats)*2.0
}

// toFloats converts a slice of ints to float64s
func toFloats(nums []int) []float64 {
	floats := make([]float64, len(nums))
	for i, n := range nums {
		floats[i] = float64(n)
	}
	return floats
}
