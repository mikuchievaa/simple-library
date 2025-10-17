package main

import (
	"fmt"
)
func main(){
	lib := &Library{}

		// Добавляем книги
	book1 := lib.AddBook("Как исключить надоедливого студента 2", "Газзаев Батраз", 1869)
	book2 := lib.AddBook("Биография Юлия Цезаря ", "Гобозов Богдан", 1866)
	book3 := lib.AddBook("1000 и 1 отмазка для прогуливания пар", "Газзаев Батраз и Гобозов Богдан", 1967)
	
	// Добавляем читателей
	reader1 := lib.AddReader("Иван", "Иванов")
	reader2 := lib.AddReader("Петр", "Петров")



// Демонстрация работы ListAllBooks - первоначальное состояние
	fmt.Println("\n--- ПЕРВОНАЧАЛЬНЫЙ КАТАЛОГ ---")
	lib.ListAllBooks()

	fmt.Println("\n=== Тест успешной выдачи ===")
	err := lib.IssueBookToReader(book1.ID, reader1.ID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println("\n=== Тест успешной выдачи второй книги ===")
	err = lib.IssueBookToReader(book2.ID, reader2.ID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	// Демонстрация работы ListAllBooks - после выдачи книг
	fmt.Println("\n--- КАТАЛОГ ПОСЛЕ ВЫДАЧИ КНИГ ---")
	lib.ListAllBooks()

	fmt.Println("\n=== Тест ошибки (книга не найдена) ===")
	err = lib.IssueBookToReader(999, reader1.ID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println("\n=== Тест ошибки (читатель не найден) ===")
	err = lib.IssueBookToReader(book3.ID, 999)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println("\n=== Тест повторной выдачи ===")
	err = lib.IssueBookToReader(book1.ID, reader2.ID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	// Возврат книги и демонстрация изменений
	fmt.Println("\n--- ТЕСТИРУЕМ ВОЗВРАТ КНИГИ ---")
	book1.ReturnBook()

	// Финальная демонстрация работы ListAllBooks
	fmt.Println("\n--- ФИНАЛЬНЫЙ КАТАЛОГ ---")
	lib.ListAllBooks()
}	



