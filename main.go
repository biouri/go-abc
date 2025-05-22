// Объявление наименования пакета
package main

// Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

// Entry function - Функция входа - функция, запускающая приложение
func main() {
	// var userKg float64 = 68 // Тип можно указать после имени переменной
	var userHeight float64 // обязательно указывается тип (по умолчанию 0.0)
	var userKg float64     // обязательно указывается тип (по умолчанию 0.0)
	fmt.Print("Рост пользователя (по умолчанию): ")
	fmt.Println(userHeight) // Рост пользователя по умолчанию 0

	// Мультистрочная Запись
	fmt.Print(`
------------------------------
Калькулятор индекса массы тела
------------------------------
Введите свой рост в сантиметрах: `)

	// Результат ввода данных будет помещен в переменную userHeight
	fmt.Scan(&userHeight) // Используется указатель для передачи параметра
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg) // Используется указатель для передачи параметра

	// := сокращенная запись создания переменных как альтернатива var
	// IMT := userKg / math.Pow(userHeight/100, IMTPower) // Возведение в степень

	// Используем функцию для вычисления ИМТ
	IMT := calculateIMT(userKg, userHeight)
	// Функция для вывода результата
	outputResult(IMT)
}

func outputResult(imt float64) {
	// Тело функции
	// Sprintf - создание строки с форматированием
	result := fmt.Sprintf("Ваш индекс массы тела (ИМТ): %.2f", imt)
	fmt.Print(result)
}

// Функция для вычисления ИМТ возвращает результат с типом float64
func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2 // Константа для возведения в степень

	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}
