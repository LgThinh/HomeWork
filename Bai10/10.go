package main

import "fmt"

func soNguyen(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Nhập số: ")
	fmt.Scanf("%d", &n)

	fmt.Println("Số nguyên tố từ 0 to", n, "là:")
	for i := 0; i <= n; i++ {
		if soNguyen(i) {
			fmt.Println(i)
		}
	}
}
