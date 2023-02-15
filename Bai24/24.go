package main

import "fmt"

func main() {
	var a, b, sumEven, sumOdd int

	fmt.Print("Nhập a: ")
	fmt.Scan(&a)

	fmt.Print("Nhập b: ")
	fmt.Scan(&b)

	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			sumEven += i
		} else {
			sumOdd += i
		}
	}

	fmt.Printf("Tổng các số chẵn từ %d đến %d là %d\n", a, b, sumEven)
	fmt.Printf("Tổng các số lẻ từ %d đến %d là %d\n", a, b, sumOdd)
}

//Trong chương trình này, chúng ta đầu tiên nhập vào hai số nguyên dương a và b.
//Sau đó, chúng ta sử dụng câu lệnh if để đảm bảo rằng a <= b.
//Tiếp theo, chúng ta sử dụng vòng lặp for để duyệt qua tất cả các số trong khoảng từ a đến b và kiểm tra xem số hiện tại là số lẻ hay số chẵn.
//Nếu số hiện tại là số chẵn, chúng ta cộng vào tổng sumEven, nếu số hiện tại là số lẻ, chúng ta cộng vào tổng sumOdd.
//Cuối cùng, chúng ta in ra tổng của tất cả các số lẻ và chẵn.
