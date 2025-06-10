package files

import (
	"fmt"
	"os"
)

// Публичная функция: Чтение файла
func ReadFile() {
	fmt.Println("Чтение файла")
}

// Публичная функция: Запись файла
func WriteFile(content string, name string) {
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
	_, err = file.WriteString(content)
	if err != nil {
		// file.Close() // Лучше использовать: defer file.Close()
		fmt.Println(err)
		return // Выход в случае ошибки
	}
	fmt.Println("Запись успешна")
	// file.Close() // Лучше использовать: defer file.Close()
}
