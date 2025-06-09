package main

import (
	"demo/app-password/account"
	"demo/app-password/files"
	"fmt"
)

func main() {
	// Запрос данных пользователя
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	// myAccount1 базовая версия Account без TimeStamp
	// Использование функции-конструктора для создания структуры
	// Функция-конструктор валидирует данные и может генерировать ошибку
	// Пароль будет сгенерирован только при отсутствии
	myAccount1, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN: " + err.Error())
		return
	}

	// Метод для вывода данных пользователя
	myAccount1.OutputPassword()

	// myAccount расширенная версия Account с TimeStamp
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN: " + err.Error())
		return
	}

	// Метод для вывода данных пользователя
	myAccount.OutputPassword()
	files.WriteFile()
	fmt.Println(myAccount)

	// Не возможно обращение к внутренним полям Account структуры
	// fmt.Println(myAccount.login)
	// fmt.Println(myAccount.Account.login)
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
