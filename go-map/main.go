package main

import "fmt"

func main() {
	// map[Тип_Ключа]:Тип_Значения{}
	// В фигурных скобках указываются значения при инициализации
	// {Ключ: Значение, ...}
	m := map[string]string{
		"DBINSK":      "https://www.dbin.sk/",
		"MebelImport": "https://www.mebelimport.su",
	}
	fmt.Println(m)

	fmt.Println(m["DBINSK"])
	// Присвоить новое значение существующему ключу (как добавление новых элементов)
	m["DBINSK"] = "https://dbin.sk"
	fmt.Println(m)

	// Добавление новых элементов
	m["Google"] = "https://google.com"
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)

	// Удаление элементов
	delete(m, "Yandex")
	// Удаление несуществующего элемнета ничего не будет
	delete(m, "Y")
	// Попытка распечатать несуществующий элемент выведет пустую строку
	fmt.Println(m["Y"])
	fmt.Println(m)

	// Итерация по map
	mapVariable := map[string]int{"a": 1, "b": 2}
	for key, value := range mapVariable {
		fmt.Println(key, value)
	}
}
