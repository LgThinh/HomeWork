package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence string
	fmt.Print("Nhập một câu: ")
	fmt.Scanf("%s", &sentence)
	words := strings.Fields(sentence)
	numWords := len(words)
	numChars := len(sentence)
	fmt.Println("Số từ trong câu là:", numWords)
	fmt.Println("Số ký tự trong câu là:", numChars)
}

//hàm fmt.Scanf được sử dụng để nhập giá trị cho biến sentence, hàm strings.Fields được sử dụng để tách câu thành từng từ
//và hai biến numWords và numChars được sử dụng để lưu số từ và ký tự trong câu.
