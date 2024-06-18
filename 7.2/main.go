package main

import (
	"fmt"
	"math"
)

func main() {
	x1 := -25
	x2 := 15
	fmt.Println(summSystem(x1, x2))
}

func summSystem(x1 int, x2 int) int {
	summ := 0
	for i := x1; i <= x2; i++ {
		if i < -5 {
			summ += 2*int(math.Abs(float64(i))) - 1
		} else if i >= -5 && i <= 5 {
			summ += i
		} else {
			summ += 2 * i
		}
	}
	return summ
}
