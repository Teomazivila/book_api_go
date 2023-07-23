package service

import (
	"bookapi/entity"
	"errors"
)

var Books []entity.Book

func InitializeBooks() {
	Books = []entity.Book{}
}
func GetAllBooks() []entity.Book {
	return Books
}

func InsertBook(book entity.Book) entity.Book {
	book.ID = uint64(len(Books) + 1)
	Books = append(Books, book)
	return book
}

func GetBook(bookID uint64) (entity.Book, error) {
	for i := 0; i < len(Books); i++ {
		if Books[i].ID == bookID {
			return Books[i], nil
		}
	}
	return entity.Book{}, errors.New("book do not exists")
}

func UpdateBook(book entity.Book) (entity.Book, error) {
	for i := 0; i < len(Books); i++ {
		if Books[i].ID == book.ID {
			Books[i] = book
			return book, nil
		}
	}
	return book, errors.New("book do not exists")
}

func DeleteBook(bookID uint64) error {
	for i := 0; i < len(Books); i++ {
		if Books[i].ID == bookID {
			Books = append(Books[:i], Books[i+1:]...)
			return nil
		}
	}
	return errors.New("book do not exists")
}
