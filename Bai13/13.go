package main

import "fmt"

func main() {
	var n int
	fmt.Print("Nhập số nguyên n: ")
	fmt.Scanf("%d", &n)
	fmt.Print("Dãy Fibonacci từ 0 đến", n, "là: ")
	a, b := 0, 1
	//dãy Fibonacci được tính dựa trên công thức F(n) = F(n-1) + F(n-2) với F(0) = 0 và F(1) = 1.
	for i := 0; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
