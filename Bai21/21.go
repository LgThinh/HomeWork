package main

import "fmt"

func main() {
	var n int
	fmt.Print("Nhập n: ")
	fmt.Scan(&n)

	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	if sum == n {
		fmt.Println(n, "là số hoàn hảo")
	} else {
		fmt.Println(n, "không phải là số hoàn hảo")
	}
}

//Trong chương trình này, chúng ta sử dụng vòng lặp for để duyệt qua các ước của n và tính tổng.
//Nếu tổng bằng n, chúng ta in ra "n là số hoàn hảo", ngược lại in ra "n không phải là số hoàn hảo".
