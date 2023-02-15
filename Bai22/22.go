package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var y int

	fmt.Print("Nhập x: ")
	fmt.Scan(&x)

	fmt.Print("Nhập y: ")
	fmt.Scan(&y)

	result := math.Pow(x, float64(y))

	fmt.Println(x, "mũ", y, "là", result)
}

//Trong chương trình này, chúng ta sử dụng hàm math.Pow(x, y) để tính xy và lưu kết quả vào biến result.
//Sau đó, chúng ta in ra kết quả của phép tính x mũ y. Chúng ta cần ép kiểu y sang kiểu float64 để hàm math.Pow hoạt động chính xác.
