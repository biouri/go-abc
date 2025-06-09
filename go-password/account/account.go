package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time // Запись с явно-именованным полем
	updatedAt time.Time // Запись с явно-именованным полем
	// Внутреннее поле Account (используется встраивание)
	// Короткая запись
	// Встраивание - аналог наследования
	Account
	// Запись с явно-именованным полем
	// Acc Account
}

// Метод структуры Account для вывода данных пользователя
// В этом методе не создается копия Acc Account т.к. используется указатель
func (acc *Account) OutputPassword() {
	fmt.Println(acc)                              // &{Login Password URL.com}
	fmt.Println(acc.login, acc.password, acc.url) // Login Password URL.com
}

// Метод структуры Account для генерации и изменения пароля
// Мутирует исходную структуру
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		loginErr := "неверный Login "
		return nil, errors.New(loginErr)
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		urlErr := "неверный URL " + err.Error()
		return nil, errors.New(urlErr)
	}
	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		// Внутреннее поле account (используется встраивание)
		Account: Account{
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
// func NewAccount(login, password, url string) *Account
// Функция-конструктор с валидацией
// 1. Если логина нет, ошибка
// 2. Если нет пароля, выполняем автогенерацию пароля
func NewAccount(login, password, urlString string) (*Account, error) {
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
	newAcc := &Account{
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	// Возвращаем созданный Account без ошибки (nil)
	return newAcc, nil
}
