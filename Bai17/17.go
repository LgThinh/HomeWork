package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Nhập số n: ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

//vòng lặp for đầu tiên được sử dụng để in ra số dòng, và vòng lặp for thứ hai được sử dụng để in ra các dấu cách.
//Vòng lặp for thứ ba được sử dụng để in ra các dấu "". Nếu số đầu tiên hoặc cuối cùng hoặc dòng cuối cùng,
//chúng ta sẽ in ra "", nếu không chúng ta sẽ in ra một khoảng trắng.