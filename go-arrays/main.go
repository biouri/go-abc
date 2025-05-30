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
