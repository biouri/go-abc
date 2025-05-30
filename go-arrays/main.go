package main

import "fmt"

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести массив
// !Вывести сумму баланса в консоль

func main() {
	tr1 := []int{1, 2, 3}
	tr2 := []int{4, 5, 6}
	tr1 = append(tr1, tr2...)
	fmt.Println(tr1)

	for index, value := range tr1 {
		fmt.Println(index, value)
	}

	// make([]ТипЭлемента, длина, ёмкость)
	// По умолчанию, если указать длину >0 будет создан массив с нулевым элементом ""
	// Итоговый массив будет иметь нулевой элемент [  1 2 3]
	tr := make([]string, 0, 2)
	fmt.Println(len(tr), cap(tr)) // 0 2
	tr = append(tr, "1")
	fmt.Println(len(tr), cap(tr)) // 1 2
	tr = append(tr, "2")
	fmt.Println(len(tr), cap(tr)) // 2 2
	tr = append(tr, "3")
	fmt.Println(len(tr), cap(tr)) // 3 4
	fmt.Println(tr)

	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)

	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f", balance)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
