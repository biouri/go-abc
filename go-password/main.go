package main

import (
	"demo/app-password/account"
	"demo/app-password/files"
	"fmt"
)

func main() {
	files.ReadFile("file.txt")

	// 1. Создать аккаунт
	// 2. Найти аккаунт
	// 3. Удалить аккаунт
	// 4. Выход
	fmt.Println("__Менеджер паролей__")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu // Выход из меню
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&variant)
	return variant
}

// Mock - функция для поиска Account
func findAccount() {

}

// Mock - функция для удаления Account
func deleteAccount() {

}

// Создать аккаунт
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
