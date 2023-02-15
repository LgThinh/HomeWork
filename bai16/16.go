package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Nhập số n: ")
	fmt.Scanf("%d", &n)
	var num int
	for i := 0; i < n; i++ {
		fmt.Print("Nhập số dương: ")
		fmt.Scanf("%d", &num)
		if num < 0 {
			fmt.Println("Số đầu vào không hợp lệ.")
			return
		}
	}
	fmt.Println("Nhập thành công.")
}

// vòng lặp for được sử dụng để nhập n số dương, và nếu một trong các số đó là số âm,
//chương trình sẽ in ra thông báo "Số đầu vào không hợp lệ." và kết thúc bằng hàm return.
