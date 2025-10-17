package main

import (
	"fmt"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int //ID читателя, который взял книгу
}

// IssueBook выдает книгу читателю
func (b *Book) IssueBook(reader *Reader) error {
	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана", b.Title)
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен", reader.FirstName, reader.LastName)
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID
	fmt.Printf("Книга '%s' была выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)
	return nil
}

// ReturnBook возвращает книгу в библиотеку
func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга '%s' и так в библиотеке", b.Title)
		return
	}
	b.IsIssued = false
	b.ReaderID = nil
	fmt.Printf("Книга '%s' возвращена в библиотеку\n", b.Title)
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен"
	} else {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, № %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

// Deactivate делает пользователя неактивным
func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
}

// Library - наша центральная структура-агрегатор
type Library struct {
	Books   []*Book
	Readers []*Reader

	//Счетчики для генерации уникальных ID
	lastBookID   int
	lastReaderID int
}

func (lib *Library) AddReader(firstName, lastName string) *Reader {
	lib.lastReaderID++

	//Создаем нового читателя
	newReader := &Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true, //Новый читатель всегда активный
	}

	//Добавляем читателя в срез
	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрирован новый читатель: %s %s \n", firstName, lastName)
	return newReader
}

// AddBook добавляет новую книгу в библиотеку
func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.lastBookID++

	//Создаем новую книгу
	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false, //Новая книга всегда в наличии
	}

	//Добавляем новую книгу в библиотеку
	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
}

// FindBookByID ищет книгу по ее уникальному ID
func (lib *Library) FindBookByID(id int) (*Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

// FindReaderByID ищет читателя по его уникальному ID
func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

// IssueBookToReader - основной публичный метод для выдачи книги
func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	// Находим книгу по ID
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err // возвращаем ошибку если книга не найдена
	}

	// Находим читателя по ID
	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err // возвращаем ошибку если читатель не найден
	}

	// Выдаем книгу читателю с помощью метода IssueBook
	err = book.IssueBook(reader)
	if err != nil {
		return err // возвращаем ошибку если выдача не удалась
	}

	return nil // успешное выполнение
}

// ListAllBooks выводит информацию о всех книгах в библиотеке
func (lib *Library) ListAllBooks() {
	fmt.Println("\n=== КАТАЛОГ ВСЕХ КНИГ В БИБЛИОТЕКЕ ===")
	if len(lib.Books) == 0 {
		fmt.Println("В библиотеке нет книг")
		return
	}
	
	for _, book := range lib.Books {
		fmt.Println(book)
	}
	fmt.Println("=== КОНЕЦ КАТАЛОГА ===")
}

func mainn() {
	lib := &Library{}

	// Добавляем книги
	book1 := lib.AddBook("Война и мир", "Лев Толстой", 1869)
	book2 := lib.AddBook("Преступление и наказание", "Федор Достоевский", 1866)
	book3 := lib.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)
	
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
