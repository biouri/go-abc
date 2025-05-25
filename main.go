// Объявление наименования пакета
package main

// Импорт нескольких пакетов
import (
	"errors"
	"fmt"
	"math"
)

// Entry function - Функция входа - функция, запускающая приложение
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // Пропустить текущую итерацию
		}
		fmt.Printf("%d\n", i)
	}

	i := 0
	for i < 10 {
		if i == 5 {
			break // Прервать выполнение цикла
		}
		fmt.Printf("%d\n", i)
		i++
	}

	// Мультистрочная Запись
	fmt.Print(`
------------------------------
Калькулятор индекса массы тела
------------------------------
`)
	// Бесконечный цикл (с выходом по условию)
	for {
		// Функция getUserInput() возвращает два параметра
		userKg, userHeight := getUserInput()

		// Если ошибку необходимо проигнорировать, используем _
		// IMT, _ := calculateIMT(userKg, userHeight)

		// Используем функцию для вычисления ИМТ
		IMT, err := calculateIMT(userKg, userHeight)

		// err == nil означает отсутствие ошибки
		// err != nil есть ошибка
		if err != nil {
			// Блок обработки ошибки
			fmt.Println(err.Error())
			fmt.Println("Не заданы параментры для расчёта")
			continue // Повторить ввод
		}

		// Функция для вывода результата
		outputResult(IMT)

		isRepeateCalculation := checkRepeatCalculation()
		// Завершение цикла при условии !isRepeateCalculation
		if !isRepeateCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	// Тело функции
	// Sprintf - создание строки с форматированием
	result := fmt.Sprintf("Ваш индекс массы тела (ИМТ): %.2f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
	/*
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
	*/
}

// При альтернативном синтаксисе возвращаемое значение (IMT float64)
// Функция для вычисления ИМТ возвращает результат float64 и ошибку
func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	const IMTPower = 2 // Константа для возведения в степень

	// Проверка корректности данных, должны быть положительные значения
	if userKg <= 0 || userHeight <= 0 {
		// Возвращает 0 и ошибку
		return 0, errors.New("PARAMS_ERROR")
	}

	// В альтернативном синтаксисе используется оператор =
	// = обозначает присваивание значения существующей переменной
	// При обычном синтаксисе := создание переменной с присваиванием
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	// Возврат результата и отсутствие ошибки
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать ещё расчёт (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
