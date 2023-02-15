package main

import "fmt"

func main() {
	var n int
	for {
		fmt.Print("Nhập n: ")
		fmt.Scanf("%d", &n)
		if n == -1 {
			fmt.Println("Vui lòng nhập lại số n khác -1")
			continue
		}
		break
	}

	for i := n; i >= -100; i-- {
		fmt.Println(i)
	}
}
