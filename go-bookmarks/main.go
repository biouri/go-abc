package main

import "fmt"

/*
Приложение сначала выдаёт меню:
- 1. Посмотреть закладки
- 2. Добавить закладку
- 3. Удалить закладку
- 4. Выход

При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/

// Объявления типов доступны на верхнем уровне
// В IDE VSCode Ctrl или Command позволяют просмотреть определение type
type bookmarkMap = map[string]string

func main() {
	// Пустой map для закладок
	// Использование Type Alias bookmarkMap вместо:
	// bookmarks := map[string]string{}
	bookmarks := bookmarkMap{}
	fmt.Println("Приложение для закладок")

Menu: // Обозначение цикла как Menu: (Label для завершения приложения)
	for {
		variant := getMenu()
		// Switch:
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		case 4:
			break Menu // Завершение приложения
			// break Switch // Завершение switch без выхода из приложения
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

// Использование Type Alias bookmarkMap вместо:
// func printBookmarks(bookmarks map[string]string)
func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

// Использование Type Alias bookmarkMap вместо:
// func addBookmark(bookmarks map[string]string) map[string]string
func addBookmark(bookmarks bookmarkMap) bookmarkMap {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

// Использование Type Alias bookmarkMap вместо:
// func deleteBookmark(bookmarks map[string]string) map[string]string
func deleteBookmark(bookmarks bookmarkMap) bookmarkMap {
	var bookmarkKeyToDelete string
	fmt.Print("Введите название: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	return bookmarks
}
