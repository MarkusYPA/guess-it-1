package ms

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
	if len(d)%2 == 0 {
		return (d[len(d)/2] + d[(len(d)/2)-1]) / 2
	} else {
		return d[len(d)/2]
	}
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

func StandardDeviation(d []float64) float64 {
	return sqrt(Variance(d))
}

// sqrt calculates the square root of a float64 according to
// the Babylonian method a.k.a. Newton's method
func sqrt(x float64) float64 {
	if x < 0 {
		return -1
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
