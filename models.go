package main

import "fmt"

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool

}
type Book struct {
	ID       int
	Title    string
	Author   string
	Year int
	IsIssued bool
	ReaderId *int
}

func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName , r.LastName, r.ID)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
	fmt.Printf("Пользователь %s %s деактивирован", r.FirstName, r.LastName)
}


func (r *Reader) Activate() {
	r.IsActive = true
	fmt.Printf("Пользователь %s %s активирован", r.FirstName, r.LastName)
}

func(r *Reader) String()string{
	status := ""
	if r.IsActive{
		status = "активен"
	}else{
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь: %s %s  (ID:%d, $s\n)", r.FirstName,r.LastName,r.ID,status)
}

func (b Book) String() string {
	return fmt.Sprintf("%s (%s, %d)", b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook(r *Reader){
	if b.IsIssued{
		fmt.Printf(" Книга %s уже выдана другому посетителю библиотеки, выберите другую\n", b.Title)
		return
	}
	// if !r.IsActive{
	// 	fmt.Printf(" Пользователь неактивен, книга %s не может быть взята\n", b.Title)
	// 	return
	// }
	b.IsIssued = true
	b.ReaderId = &r.ID
	fmt.Printf("Книга %s выдана: %s %s\n", b.Title, r.FirstName, r.LastName )
}

func (b *Book) ReturnBook(){
	if !b.IsIssued{
		fmt.Printf("Книга %s ни у кого не находится\n", b.Title)
		return
	}
	b.IsIssued = false
	b.ReaderId = nil 
	fmt.Printf("Книга %s возвращена\n", b.Title)
}