package main

import (
	"fmt"
)

func main() {
	// Объявление массива из 3-элементов
	// В обоих случаях массив будет проинициализирован нулевыми значениями

	// Объявление переменной с явным указанием типа:
	// Все элементы массива инициализируются нулевыми значениями
	var transactionsA [3]int // [0 0 0] Работает везде (пакет/функция)

	// Это короткое объявление переменной с инициализацией пустыми значениями
	transactionsB := [3]int{} // [0 0 0] Работает только внутри функций

	// Инициализировать с предопределённым набором элементов
	transactions := [6]int{1, 2, 3, 4, 5, 6}
	banks := [2]string{"Тинькофф", "Альфа"}
	wallets := [2]string{} // Объявить как переменную без присвоения
	wallets[0] = "Черный"  // Записать определенный элемент по индексу

	fmt.Println(transactionsA)   // [0 0 0]
	fmt.Println(transactionsB)   // [0 0 0]
	fmt.Println(transactions)    // [1 2 3 4 5 6]
	fmt.Println(transactions[1]) // 2 Прочитать определенный элемент по индексу
	fmt.Println(banks)           // [Тинькофф Альфа]
	fmt.Println(wallets)         // [Черный ]

	copy := transactions            // Это полная независимая копия массива
	copy[1] = 100                   // Изменение не повлияет на исходный массив
	fmt.Println(transactions, copy) // [1 2 3 4 5 6] [1 100 3 4 5 6]

	// Slice
	patrial := transactions[1:] // Часть от 1 до последнего
	fmt.Println(patrial)        // [2 3 4 5 6]

	// Slice ссылается на исходный массив
	// Любое изменение slice приводит к изменению исходного массива
	transactionsPartial := transactions[1:5]          // slice от array
	fmt.Println(transactionsPartial)                  // [2 3 4 5]
	transactionsNewPartial := transactionsPartial[:1] // slice от slice
	fmt.Println(transactionsNewPartial)               // [2]
	transactionsNewPartial[0] = 30                    // изменит slice и array

	fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))

	transactionsNewPartial = transactionsNewPartial[0:4] // это уже другой slice

	fmt.Println(transactions)           // [1 30 3 4 5 6]
	fmt.Println(transactionsNewPartial) // [30 3 4 5] часть массива transactions

	// Length (Длина): Актуальное количество элементов в слайсе.
	// Capacity (Вместимость): Максимальное количество элементов,
	// которое слайс может содержать, основываясь на размере исходного массива.

	fmt.Println(len(transactionsPartial), cap(transactionsPartial))       // 4 5
	fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial)) // 4 5

	// Динамическим может быть только slice
	// slice не имеет фиксированной длины изначально
	// Под капотом создается массив на который ссылается slice
	transactionSlice := []int{0, 20, 35}
	temp := transactionSlice
	// append() нельзя использовать для array, но можно использовать для slice
	transactionSlice = append(transactionSlice, 100) // append возвращает новый slice

	fmt.Println(temp)                 // [0 20 35] temp ссылается на другой массив
	fmt.Println(transactionSlice[1:]) // [20 35 100] новй slice без 0-го элемента
}
