package dungeongenerator

import (
	"fmt"
	"math"
	"math/rand"
)

func RandInt(min, max int) int {
	if max < min {
		fmt.Println("swapping")
		max, min = min, max
	}
	n := max - min
	if n < 0 {
		n = int(math.Abs(float64(n)))
		return min + rand.Intn(n)
	} else if n == 0 {
		return min + rand.Intn(n+1)
	}
	return min + rand.Intn(max-min)
}
