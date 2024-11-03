package main

import (
	"fmt"
	"math"
)

func mysqrt(x int) int {
	temp := float64(x)
	return int(math.Sqrt(temp))
}

func main() {
	fmt.Println(mysqrt(2))
	fmt.Println(mysqrt(8))
}
