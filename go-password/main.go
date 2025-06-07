package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

// Метод структуры account для вывода данных пользователя
// В этом методе не создается копия acc account т.к. используется указатель
func (acc *account) outputPassword() {
	fmt.Println(acc)                              // &{Login Password URL.com}
	fmt.Println(acc.login, acc.password, acc.url) // Login Password URL.com
}

// Метод структуры account для генерации и изменения пароля
// Мутирует исходную структуру
func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	str := []rune("Привет! 😊")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}

	str2 := "Привет! 😊"
	for _, ch := range str2 {
		fmt.Println(ch, string(ch))
	}

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

	// Генерация случайного числа от 0 до 9
	// IntN(n int) int
	// IntN returns, as an int, a pseudo-random number
	// in the half-open interval [0,n) from the default Source.
	// It panics if n <= 0.
	fmt.Println(rand.IntN(10))

	// Функция - Генерации случайного пароля из 12 символов
	fmt.Println(generatePassword(12))

	// Запрос данных пользователя
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	// Рекомендуемый (основной) сопособ создания экземпляра структуры
	// Порядок может быть любой, можно часть значений не определять
	// Не заданные значения будут иметь значения по умолчанию
	myAccount := account{
		password: password,
		url:      url,
		login:    login,
	}

	// Альтернативный способ объявления переменной для структуры
	// В данном случае важен порядок следования значений структуры
	// В этом примере нарушен порядок, первым должен быть login
	// Поле login будет иметь значение переменной password
	account1 := account{
		password,
		"", // Не определенные значения нужно указывать явно
		url,
	}

	// Создание пустой структуры для последующего заполнения через методы
	account2 := account{}

	// До изменения пароля myAccount
	fmt.Println(myAccount, account1, account2)
	// Метод для вывода данных пользователя
	myAccount.outputPassword()

	// Метод для генерации и изменения пароля пользователя
	myAccount.generatePassword(12)
	// Метод для вывода данных пользователя
	myAccount.outputPassword()
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

// Ввод данных с консоли
func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

// Функция - Генератор строки из n случайных символов
func generatePassword(n int) string {
	// Slice определенной длины n
	res := make([]rune, n)
	for i := range res {
		// Берем случайный элемент из массива допустимых символов
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	// Преобразование результата в строку
	return string(res)
}
