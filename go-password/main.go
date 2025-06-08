package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time // Запись с явно-именованным полем
	updatedAt time.Time // Запись с явно-именованным полем
	// Внутреннее поле account (используется встраивание)
	// Короткая запись
	// Встраивание - аналог наследования
	account
	// Запись с явно-именованным полем
	// acc account
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

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		loginErr := "неверный Login "
		return nil, errors.New(loginErr)
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		urlErr := "неверный URL " + err.Error()
		return nil, errors.New(urlErr)
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		// Внутреннее поле account (используется встраивание)
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

// Сигнатура функции-конструктора без валидации
// func newAccount(login, password, url string) *account
// Функция-конструктор с валидацией
// 1. Если логина нет, ошибка
// 2. Если нет пароля, выполняем автогенерацию пароля
func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		// В Go принято, что строки ошибок:
		// Начинаются со строчной буквы (кроме имён собственных или аббревиатур)
		// Не заканчиваются точкой (т.к. ошибки часто оборачиваются в другие сообщения)
		loginErr := "неверный Login "
		return nil, errors.New(loginErr)
	}
	// Перед возвратом структуры можно выполнить валидацию
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		// Собственная ошибка создается если нужно показать пользователю
		// новое сообщение на понятном родном языке, также можно
		// добавить оригинальное сообщение об ошибке
		// Строчная буква, без точки
		urlErr := "неверный URL " + err.Error()
		return nil, errors.New(urlErr)
	}
	newAcc := &account{
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	// Возвращаем созданный account без ошибки (nil)
	return newAcc, nil
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
	// login := ""
	// password := ""
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	// Рекомендуемый (основной) сопособ создания экземпляра структуры
	// Порядок может быть любой, можно часть значений не определять
	// Не заданные значения будут иметь значения по умолчанию
	// myAccount := account{
	// 	password: password,
	// 	url:      url,
	// 	login:    login,
	// }

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

	fmt.Println(account1, account2)

	myAccount1, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN: " + err.Error())
		return
	}

	// Метод для вывода данных пользователя
	myAccount1.outputPassword()

	// Использование функции-конструктора для создания структуры
	// Функция-конструктор newAccount валидирует данные и может генерировать ошибку
	// Пароль будет сгенерирован только при отсутствии
	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN: " + err.Error())
		return
	}

	// Метод для генерации и изменения пароля пользователя
	// Теперь интегрирован в Функцию-конструктор
	// myAccount.generatePassword(12)

	// Метод для вывода данных пользователя
	myAccount.outputPassword()

	// Поля с типом time.Time также являются структурами
	// &{{13981831876230078340 20863322501 0x50f880} {13981831876230078340 20863322501 0x50f880} {Login Password http://url.com}}
	fmt.Println(myAccount)
	// Доступ к явно-именованным полям
	fmt.Println(myAccount.createdAt) // 2025-06-08 12:05:56.2925118 +0300 MSK m=+15.578619301
	fmt.Println(myAccount.updatedAt) // 2025-06-08 12:05:56.2925118 +0300 MSK m=+15.578619301

	// Возможно обращение к внутренним полям account структуры двумя способами
	fmt.Println(myAccount.login)
	fmt.Println(myAccount.account.login)
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
	fmt.Scanln(&res) // Позволяет ввести пустую строку
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
