package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой")

	//1. Создаем экземпляр библиотеки
	myLibrary := &Library{} //Пустая библиотека готова к работе

	fmt.Println("Наполняем библиотеку")
	//2. Добавляем читателей
	myLibrary.AddReader("Агунда", "Кокойти")
	myLibrary.AddReader("Сергей", "Меняйло")

	//3. Добавляем книги
	myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

	fmt.Println("\n---Библиотека готова к работе---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))

	//Модуль 16. Практикум
	fmt.Println("---Тестируем выдачу книг---")
	//Выдаем книгу 1 читателю 1
	fmt.Println("Попытка выдать книгу 1 читателю 1")
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}

	//Попытка выдать ту же книгу еще раз
	fmt.Println("Поптка выдать уже выданную книгу")
	err = myLibrary.IssueBookToReader(1, 2)
	if err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}

	fmt.Println("Попытка выдать несуществующую книгу")
	err = myLibrary.IssueBookToReader(99, 1)
	if err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}

	fmt.Println("Попытка выдать книгу несуществующему читателю")
	err = myLibrary.IssueBookToReader(2, 99)
	if err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга успешно выдана")
	}

	fmt.Println()

	//Смотрим все книги в библиотеке
	myLibrary.ListAllBooks()

	fmt.Println("Поиск порта")
	config := map[string]string{
		"PORT": "456",
	}

	port, err := GetPortFromConfig(config)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(port)
	}

	config = map[string]string{
		"tort": "medivik",
	}
	port, err = GetPortFromConfig(config)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(port)
	}
}
