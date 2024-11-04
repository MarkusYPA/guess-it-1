package guessing

import (
	"math"

	"01.gritlab.ax/git/mamberla/guess-it-1/ms"
)

// MedAndSD returns the median of the data plus and minus a multiple of the standard deviation
func MedAndSD(nums []float64, mult float64) [2]int {
	rng := [2]int{}
	sd := ms.StandardDeviation(nums) * mult

	if len(nums) == 1 {
		rng[0] = ms.RoundToInt(nums[0]) - ms.RoundToInt((ms.Abs(nums[0]))/2.0)
		rng[1] = ms.RoundToInt(nums[0]) + ms.RoundToInt((ms.Abs(nums[0]))/2.0)
	} else {
		rng[0] = ms.RoundToInt(ms.Median(nums) - sd)
		rng[1] = ms.RoundToInt(ms.Median(nums) + sd)
	}

	return rng
}

// AvgAndSD returns the mean of the data plus and minus a multiple of the standard deviation
func AvgAndSD(nums []float64, mult float64) [2]int {
	rng := [2]int{}
	sd := ms.Abs(ms.StandardDeviation(nums)) * mult

	if len(nums) == 1 {
		rng[0] = ms.RoundToInt(nums[0]) - ms.RoundToInt((ms.Abs(nums[0]))/2.0)
		rng[1] = ms.RoundToInt(nums[0]) + ms.RoundToInt((ms.Abs(nums[0]))/2.0)
	} else {
		rng[0] = ms.RoundToInt(ms.Average(nums) - sd)
		rng[1] = ms.RoundToInt(ms.Average(nums) + sd)
	}

	return rng
}

// Box returns 1st and 3rd quartile cutting points, often seen in box plots
func Box(nums []float64) [2]int {
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

// Silly guess returns a range simply using the two previous data points
func SillyGuess(nums []int) [2]int {
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
		if nums[len(nums)-2] < nums[len(nums)-1] {
			rng[0] = nums[len(nums)-2]
			rng[1] = nums[len(nums)-1]
		} else {
			rng[0] = nums[len(nums)-1]
			rng[1] = nums[len(nums)-2]
		}
	}

	return rng
}

// HardRange returns a static range of known frequent numbers
func HardRange() [2]int {
	//return [2]int{113, 113}
	return [2]int{112, 113}
	//return [2]int{112, 114}
	//return [2]int{112, 115}
	//return [2]int{112, 116}
}

// MedToMean returns a range from median to mean (or vice versa) plus minus a multiple of the standard deviation
func MedToMean(nums []float64, mult float64) [2]int {
	mean := int(math.Round(ms.Average(nums)))
	medi := int(math.Round(ms.Median(nums)))
	sd := int(math.Round(ms.StandardDeviation(nums) * mult))

	if medi < mean {
		return [2]int{medi - sd, mean + sd}
	} else {
		return [2]int{mean - sd, medi + sd}
	}
}
