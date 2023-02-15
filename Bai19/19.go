package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Nhập x: ")
	fmt.Scanf("%d", &x)
	fmt.Print("Nhập y: ")
	fmt.Scanf("%d", &y)

	sum := 0
	for i := x; i <= y; i++ {
		sum += i * i
	}
	fmt.Printf("Tổng bình phương các số từ %d đến %d là: %d\n", x, y, sum)
}
