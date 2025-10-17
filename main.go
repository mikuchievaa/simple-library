package main

import (
	"fmt"
)



func main(){
	myLibrary := Library{
		Books: []*Book{
			{ID:1,Title: "Как исключить надоедливого студента 2",Author:"Газзаев Батраз",Year:3892,IsIssued:false},
			{ID:2,Title: "Биография Юлия Цезаря ",Author:"Гобозов Богдан",Year:2025,IsIssued:false},
			{ID:3,Title: "1000 и 1 отмазка для прогуливания пар",Author:"Газзаев Батраз и Гобозов Богдан",Year:2002,IsIssued:false},

		},
		Readers: []*Reader{
			{ID:1,FirstName: "Илона",LastName: "Валиева",IsActive: false},
			{ID:2,FirstName: "Милана",LastName: "Кучиева",IsActive: false},
			{ID:3,FirstName: "Марлен",LastName: "Албегов",IsActive: false},

		},
	}



	fmt.Println("---Тестируем выдачу книг---")
	//Выдаем книгу 1 читателю 1
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
	fmt.Println("Ошибка выдачи", err)
	}
	//Проверить статус книги после выдачи
	book, _ := myLibrary.FindBookByID(1)
	if book != nil {
	fmt.Println("Статус книги после выдачи:", book)
	}
	//Попытка выдать несуществующую книгу
	err = myLibrary.IssueBookToReader(99, 1)
	if err != nil {
	fmt.Println("Ожидаемая ошибка:", err)
	}
}	

