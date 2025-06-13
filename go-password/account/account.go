package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	// Поля записываются с заглавной, чтобы экспортировались
	// Добавлена метаинформация
	// Пример структурных тегов для полей структуры: `json:"login"` ...
	// имена записаны с маленькой буквы (соответствует стилю именования в JSON)
	// Если не указать метаинформацию, будут использованы имена из структуры.
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Метод структуры Account для вывода данных пользователя
// В этом методе не создается копия Acc Account т.к. используется указатель
func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login + " " + acc.Password + " " + acc.Url) // цвет: Cyan
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

// Метод преобразования Account в byte-массив
// Имя метода ToBytes с заглавной т.к. он экспортируется
// type byte = uint8 (byte является алиасом к типу uint8)
func (acc *Account) ToBytes() ([]byte, error) {
	// Аналогичная функция преобразования с отступами: json.MarshalIndent()
	// Возвращается: массив_байтов file, ошибка err
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Метод структуры Account для генерации и изменения пароля
// Мутирует исходную структуру
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

// Функция-конструктор с валидацией
// 1. Если нет логина, ошибка
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}

	// Получить метатеги поля login в Runtime
	// reflect библиотека позволяет работать с типами в Runtime
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("Login")
	// Получить метаинформацию о поле login
	fmt.Println("Meta Information")
	fmt.Println(string(field.Tag)) // json:"login"

	if password == "" {
		newAcc.generatePassword(12)
	}
	// Возвращаем созданный Account без ошибки (nil)
	return newAcc, nil
}
