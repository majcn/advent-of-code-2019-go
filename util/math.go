package util

import "math"

func Min(s []int) (rc int) {
	rc = s[0]

	for i := 1; i < len(s); i++ {
		if s[i] < rc {
			rc = s[i]
		}
	}

	return
}

func Max(s []int) (rc int) {
	rc = s[0]

	for i := 1; i < len(s); i++ {
		if s[i] > rc {
			rc = s[i]
		}
	}

	return
}

func Sum(s []int) (rc int) {
	for i := 0; i < len(s); i++ {
		rc += s[i]
	}

	return
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func PowInt(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
