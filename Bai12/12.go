package main

import "fmt"

func main() {
	var n int
	fmt.Print("Nhập số nguyên dương n: ")
	fmt.Scanf("%d", &n)

	fmt.Print(n, " phần tử đầu của dãy Fibonacci: ")
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
