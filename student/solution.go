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
func medAndSD(nums []float64, mult float64) [2]int {
	rng := [2]int{}

	if len(nums) == 1 {
		rng[0] = ms.RoundToInt(nums[0]) - ms.RoundToInt((ms.Abs(nums[0]))/2.0)
		rng[1] = ms.RoundToInt(nums[0]) + ms.RoundToInt((ms.Abs(nums[0]))/2.0)
	} else {
		rng[0] = ms.RoundToInt(ms.Median(nums) - ms.Abs(ms.StandardDeviation(nums))*mult)
		rng[1] = ms.RoundToInt(ms.Median(nums) + ms.Abs(ms.StandardDeviation(nums))*mult)
	}

	return rng
}

// avgAndSD returns the mean of the data plus and minus the standard deviation
func avgAndSD(nums []float64, mult float64) [2]int {
	rng := [2]int{}

	if len(nums) == 1 {
		rng[0] = ms.RoundToInt(nums[0]) - ms.RoundToInt((ms.Abs(nums[0]))/2.0)
		rng[1] = ms.RoundToInt(nums[0]) + ms.RoundToInt((ms.Abs(nums[0]))/2.0)
	} else {
		rng[0] = ms.RoundToInt(ms.Average(nums) - ms.Abs(ms.StandardDeviation(nums))*mult)
		rng[1] = ms.RoundToInt(ms.Average(nums) + ms.Abs(ms.StandardDeviation(nums))*mult)
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

func sillyGuess(nums []int) [2]int {
	rng := [2]int{}

	if len(nums) == 1 {
		rng[0] = nums[0]
		rng[1] = nums[0]
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			rng[0] = nums[0]
			rng[1] = nums[1]
		} else {
			rng[0] = nums[1]
			rng[1] = nums[0]
		}
	} else {
		if nums[len(nums)-3] < nums[len(nums)-1] {
			rng[0] = nums[len(nums)-3]
			rng[1] = nums[len(nums)-1]
		} else {
			rng[0] = nums[len(nums)-1]
			rng[1] = nums[len(nums)-3]
		}
	}

	return rng
}

// guessNextRange makes a guess about in which range the next number will be
func guessNextRange(nums []int, mult float64) (int, int) {

	if len(nums) == 0 {
		return 0, 0
	}

	numsF := []float64{}
	for _, n := range nums {
		numsF = append(numsF, float64(n))
	}

	//rng := box(numsF)
	//rng := avgAndSD(numsF, mult)
	rng := medAndSD(numsF, mult)
	//rng := [2]int{nums[len(nums)-1], nums[len(nums)-1]}
	//rng := sillyGuess(nums)

	return rng[0], rng[1]
}

func main() {
	pointsOn := flag.Bool("points", false, "Display the points at the end or not")
	flag.Parse()

	mult := 0.67
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

			if seen > 5 && isOutlier(num, numbers) {
				fmt.Println(guess)
				continue
			}

			// Only ever deal with data sets maximum 100 long
			if seen < 100 {
				seen++
				numbers = append(numbers, num)
			} else {
				numbers = append(numbers[1:], num)
			}

			fmt.Printf("%s  %v %v\n", guess, num, points)

			r1, r2 = guessNextRange(numbers, mult)
			guess = fmt.Sprintf("%v %v", r1, r2)

			//fmt.Println(100, 200)
		}
	}
	if *pointsOn {
		fmt.Printf("\n%v points aquired\n", points)
	}
}

func getPoints(num, r1, r2 int) int {
	points := 0.0
	rangeWidth := r2 - r1 + 1
	if num >= r1 && num <= r2 {
		points += 800 / float64(rangeWidth)
	}
	return int(math.Round(points))
}

func isOutlier(n int, nums []int) bool {
	floats := toFloats(nums)
	//fmt.Println("Testing outlier, this greater than that?", ms.Abs(float64(n)-ms.Average(floats)), ms.Variance(floats)*2.0, ms.Abs(float64(n)-ms.Average(floats)) > ms.Variance(floats)*2.0)
	return ms.Abs(float64(n)-ms.Average(floats)) > ms.Variance(floats)*2.0
}

func toFloats(nums []int) []float64 {
	floats := make([]float64, len(nums))
	for i, n := range nums {
		floats[i] = float64(n)
	}
	//fmt.Println("floats:", floats)
	return floats
}
