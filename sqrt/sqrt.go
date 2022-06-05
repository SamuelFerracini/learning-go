package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-float64(0)) < 1e-10 {
			break
		}

	}

	return z
}

func main() {
	fmt.Println(Sqrt(20))
}
