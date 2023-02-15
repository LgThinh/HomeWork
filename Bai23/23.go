package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Nhập n: ")
	fmt.Scan(&n)

	isPrime := true

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if n < 2 {
		isPrime = false
	}

	if isPrime {
		fmt.Println(n, "là số nguyên tố")
	} else {
		fmt.Println(n, "không phải là số nguyên tố")
	}
}

//Trong chương trình này, chúng ta sử dụng vòng lặp for để duyệt qua các số nguyên tố nhỏ hơn hoặc bằng căn bậc hai của n.
//Nếu n chia hết cho bất kỳ số nguyên tố nào trong đó, chúng ta gán isPrime = false và thoát vòng lặp.
//Nếu n < 2, chúng ta cũng gán isPrime = false vì 1 và các số âm không được coi là số nguyên tố.
//Cuối cùng, chúng ta kiểm tra giá trị của isPrime và in ra kết quả tương ứng.