package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
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

	fmt.Println(myAccount, account1, account2)

	// Моковая функция для вывода данных пользователя
	// Входной параметр - указатель на структуру
	outputPassword(&myAccount)
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

// Моковая функция для вывода данных пользователя
func outputPassword(acc *account) {
	fmt.Println(acc)                              // &{Login Password URL.com}
	fmt.Println(acc.login, acc.password, acc.url) // Login Password URL.com
	// Запись тождественна т.к. dereference выполняется для структур автоматически
	// Неявный dereference - получение значения по адресу в памяти
	fmt.Println((*acc).login, (*acc).password, (*acc).url) // Login Password URL.com
}
