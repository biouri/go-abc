package main

import "fmt"

func main() {
	a := 5
	// var pointerA *int - это указатель на переменную int
	pointerA := &a // Создание указателя на переменную a
	res := double(a)
	fmt.Println(pointerA) // 0xc00000a0e8
	fmt.Println(res)      // 10
}

func double(num int) int {
	return num * 2
}
