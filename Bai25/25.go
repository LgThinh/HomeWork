package main

import "fmt"

func main() {
	var n int
	var sum float64

	fmt.Print("Nhập n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		sum += float64(i) * float64(i+1)
	}

	fmt.Printf("Giá trị của biểu thức là %f", sum)
}

//Trong chương trình này, chúng ta đầu tiên nhập vào số nguyên dương n từ bàn phím.
//Sau đó, chúng ta sử dụng vòng lặp for để duyệt qua các số từ 1 đến n, và tính tổng của từng phần tử trong biểu thức bằng cách nhân i với (i+1) và cộng vào biến sum.
//Cuối cùng, chúng ta in ra giá trị của biểu thức S.