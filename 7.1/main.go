package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	fmt.Println(CountOfDigits(n) * CountOfDigits(m))
}

func CountOfDigits(n int) int {
	counter := 0
	for n > 0 {
		counter++
		n /= 10
	}
	return counter
}
