package domain

import (
	
	"fmt"
	// "strings"
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
	return nil
}

// ReturnBook возвращает книгу в библиотеку
func (b *Book) ReturnBook() error{
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
		
	}
	b.IsIssued = false
	b.ReaderID = nil
	return nil
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
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

func (r *Reader) Activate() {
	r.IsActive = true
}






