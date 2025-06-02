package main

import "fmt"

func main() {
	a := 5
	// var pointerA *int - это указатель на переменную int
	pointerA := &a // Создание указателя на переменную a
	// Оператор * dereference: значение по адресу памяти
	fmt.Println(*pointerA) // 5
	double(&a)
	fmt.Println(a) // 10

	m := [4]int{1, 2, 3, 4}
	reverseOK(&m) // Меняет порядок элементов на обратный
	fmt.Println(m)
}

// Функция double принимает ссылку на int и ничего не возвращает
// В результате происходит мутирование (изменение) исходного объекта
func double(num *int) {
	// Мутирование исходного объекта
	*num = *num * 2 // *num - * получить значение из адреса в памяти
}

// Входной параметр функции - ссылка на массив из четырех элементов
// Функция работает, но она неэффективна, поскольку
// создается копия массива при выполнении range *arr.
// Достаточно пройтись только до середины и выполнять обмен элементов
// на каждом шаге попарно.
func reverse(arr *[4]int) {
	for index, value := range *arr {
		fmt.Println("old:", *arr)
		(*arr)[len(arr)-1-index] = value
		fmt.Println(index, ":", value, "move to:", len(arr)-1-index, "new:", *arr)
	}
}

// Чтобы реверсировать массив, достаточно пройтись только до середины
// и на каждой итерации менять местами зеркальные элементы
func reverseOK(arr *[4]int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
		fmt.Println("Swapped", i, "and", j, ":", *arr)
	}
}
