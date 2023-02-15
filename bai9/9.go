package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Nhập một dãy số: ")
	fmt.Scan(&name)
	nameRev := reverse(name)
	fmt.Println("Dãy số đã nhập là: ", name)
	fmt.Println("Dãy số đảo ngược là: ", nameRev)

}
func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
