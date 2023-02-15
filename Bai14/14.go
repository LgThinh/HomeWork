package main

import "fmt"

func main() {
	var x int
	fmt.Print("Nhập số nguyên dương x: ")
	fmt.Scanf("%d", &x)
	result := 1
	for i := 1; i <= x; i++ {
		result *= i
	}
	fmt.Println("Giai thừa của", x, "là", result)
}
