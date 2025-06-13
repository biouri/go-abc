package files

import (
	"fmt"
	"os"
)

// Публичная функция: Чтение файла
func ReadFile(filename string) {
	fmt.Println("Чтение файла")
	// file, err := os.Open(filename) // Побайтовое чтение (по порциям)
	// ioutil.ReadFile(filename) // Deprecated старый нерекомендуемый

	// os.ReadFile открыть и получить данные в виде массива байтов
	data, err := os.ReadFile(filename) // Открыть и полностью прочитать
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)         // Массив байтов
	fmt.Println(string(data)) // Преобразовать массив байтов в строку
}

// Публичная функция: Запись файла
func WriteBytesToFile(content []byte, name string) {
	fmt.Println("Запись файла")
	file, err := os.Create(name) // (пере)Создание файла
	if err != nil {
		fmt.Println(err)
		return // Выход в случае ошибки
	}

	// Предпочтительный способ закрытия файлов
	// defer - отложить выполнение некоторых операций до момента очистки стек-фрейма
	defer file.Close()

	// Возвращается len, err (кол-во записанных байт можно не пропустить)
	// Вместо WriteString используем Write
	_, err = file.Write(content)
	if err != nil {
		// file.Close() // Лучше использовать: defer file.Close()
		fmt.Println(err)
		return // Выход в случае ошибки
	}
	fmt.Println("Запись успешна")
	// file.Close() // Лучше использовать: defer file.Close()
}
