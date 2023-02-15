package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Nhập n: ")
	fmt.Scanf("%d", &n)

	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1 / math.Pow(float64(i), 3)
	}
	fmt.Printf("S=%.5f\n", sum)
}

//biểu thức math.Pow(x, y) trong Go tính x^y.
