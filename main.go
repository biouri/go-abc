// Объявление наименования пакета
package main

// Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

// Entry function - Функция входа - функция, запускающая приложение
func main() {
	// Мультистрочная Запись
	fmt.Print(`
------------------------------
Калькулятор индекса массы тела
------------------------------
`)
	// Функция getUserInput() возвращает два параметра
	userKg, userHeight := getUserInput()

	// Используем функцию для вычисления ИМТ
	IMT := calculateIMT(userKg, userHeight)

	// Функция для вывода результата
	outputResult(IMT)

	if IMT < 16 {
		fmt.Println("У вас сильный дефицит массы тела")
	} else if IMT < 18.5 {
		fmt.Println("У вас дефицит массы тела")
	} else if IMT < 25 {
		fmt.Println("У вас нормальный вес")
	} else if IMT < 30 {
		fmt.Println("У вас избыточный вес")
	} else {
		fmt.Println("У вас степень ожирения")
	}
}

func outputResult(imt float64) {
	// Тело функции
	// Sprintf - создание строки с форматированием
	result := fmt.Sprintf("Ваш индекс массы тела (ИМТ): %.2f", imt)
	fmt.Println(result)
}

// Альтернативный синтаксис возвращаемое значение (IMT float64)
// Функция для вычисления ИМТ возвращает результат с типом float64
func calculateIMT(userKg float64, userHeight float64) (IMT float64) {
	const IMTPower = 2 // Константа для возведения в степень

	// В альтернативном синтаксисе используется оператор =
	// = обозначает присваивание значения существующей переменной
	IMT = userKg / math.Pow(userHeight/100, IMTPower)
	// В альтернативном синтаксисе не указывается имя переменной в return
	return
}

// Функция с двумя возвращаемыми параметрами (float64, float64)
func getUserInput() (float64, float64) {
	var userHeight float64 // обязательно указывается тип (по умолчанию 0.0)
	var userKg float64     // обязательно указывается тип (по умолчанию 0.0)
	fmt.Print("Введите свой рост в сантиментрах: ")

	// Результат ввода данных будет помещен в переменную userHeight
	fmt.Scan(&userHeight) // Используется указатель для передачи параметра
	fmt.Print("Введите свой вес: ")

	// Результат ввода данных будет помещен в переменную userKg
	fmt.Scan(&userKg) // Используется указатель для передачи параметра

	// Возврат из функции двух параметров
	return userKg, userHeight
}
