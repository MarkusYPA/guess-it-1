package ms

import "sort"

// roundToInt rounds a floa64 to an integer
func RoundToInt(f float64) int {
	diff := f - float64(int(f))

	if diff >= 0.5 {
		return int(f) + 1
	} else if diff <= -0.5 {
		return int(f) - 1
	} else {
		return int(f)
	}
}

// average returns the mean of a slice of float64s
func Average(d []float64) float64 {
	sum := 0.0
	for _, v := range d {
		sum += v
	}
	return sum / float64(len(d))
}

// median returns the median of a slice of float64s
func Median(d []float64) float64 {
	//return medianWithSorted(bubSort(d))
	return medianWithSorted(sortImp(d))
}

func medianWithSorted(d []float64) float64 {
	if len(d) == 0 {
		return 0
	}
	if len(d)%2 == 0 {
		return (d[len(d)/2] + d[(len(d)/2)-1]) / 2
	} else {
		return d[len(d)/2]
	}
}

// Quarters returns the minimum, 1st quartile, median, 3rd quartile and maximum values
func Quarters(d []float64) [5]float64 {
	qs := [5]float64{}

	//dSorted := bubSort(d)
	dSorted := sortImp(d)
	ln := len(dSorted)

	qs[0] = dSorted[0]
	//qs[1] = medianWithSorted(dSorted[:ln/2])
	qs[2] = medianWithSorted(dSorted)
	//qs[3] = medianWithSorted(dSorted[(ln+1)/2:])
	qs[4] = dSorted[ln-1]

	qs[1] = medianWithSorted(dSorted[(ln+1)/4 : ln/2])      // 3/8 point instead
	qs[3] = medianWithSorted(dSorted[(ln+1)/2 : ln-(ln/4)]) // 5/8 point instead

	return qs
}

// bubSort is a bubble sort function that arranges a slice of float64s from smallest to largest
func bubSort(d []float64) []float64 {
	for i := 0; i < len(d)-1; i++ {
		for j := i + 1; j < len(d); j++ {
			if d[i] > d[j] {
				d[i], d[j] = d[j], d[i]
			}
		}
	}
	return d
}

// sortImp imports a sorting algorithm and sorts a slice of float64s with it
func sortImp(d []float64) []float64 {
	dSorted := make([]float64, len(d))
	copy(dSorted, d)
	sort.Float64s(dSorted)
	return dSorted
}

// variance returns the variance of a slice of float64s
func Variance(d []float64) float64 {
	sumOfSqOfDiff := 0.0
	avg := Average(d)
	for _, f := range d {
		sumOfSqOfDiff += (f - avg) * (f - avg)
	}
	return sumOfSqOfDiff / float64(len(d))
}

// StandardDeviation returns the square root of Variance()
func StandardDeviation(d []float64) float64 {
	return sqrt(Variance(d))
}

// sqrt calculates the square root of a float64 according to
// the Babylonian method a.k.a. Newton's method
func sqrt(x float64) float64 {
	if x < 0 {
		return -1.0
	}
	if x == 0 {
		return 0.0
	}

	// start with a guess
	z := x
	const tolerance = 1e-10 // how precise you want the result to be
	for {
		nextZ := 0.5 * (z + x/z)
		if Abs(z-nextZ) < tolerance { // Stop when the change is smaller than the tolerance
			break
		}
		z = nextZ
	}
	return z
}

// abs calculates the absolute value of a float64
func Abs(f float64) float64 {
	if f < 0 {
		f *= -1
	}
	return f
}
