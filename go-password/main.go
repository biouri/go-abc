package main

import (
	"demo/app-password/account"
	"demo/app-password/files"
	"fmt"
)

func main() {
	files.ReadFile("file.txt")

	createAccount()
}

func createAccount() {
	// Запрос данных пользователя
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN: " + err.Error())
		return
	}

	// Метод для вывода данных пользователя в консоль
	myAccount.OutputPassword()
	fmt.Println(myAccount)

	// Сохранение структуры в файл
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteBytesToFile(file, "data.json")
}

// Ввод данных с консоли
func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res) // Позволяет ввести пустую строку
	return res
}

// Функция - Генератор строки из n случайных символов
// Может быть и функция и метод с одним именем
// Метод работает только для определенной структуры
// Функция более универсальна и может работать в любом месте
// func generatePassword(n int) string {
// 	// Slice определенной длины n
// 	res := make([]rune, n)
// 	for i := range res {
// 		// Берем случайный элемент из массива допустимых символов
// 		res[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	// Преобразование результата в строку
// 	return string(res)
// }
