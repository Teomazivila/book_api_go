package service

import (
	"bookapi/entity"
	"bookapi/repository"
	"errors"
)

func GetAllBooks() []entity.Book {
	return repository.GetAllBooks()
}

func InsertBook(book entity.Book) entity.Book {
	book.UserID = 2
	book = repository.InsertBook(book)

	return book
}

func GetBook(bookID uint64) (entity.Book, error) {
	if book, err := repository.GetBook(bookID); err == nil {
		return book, nil
	}
	return entity.Book{}, errors.New("book do not exists")
}

func UpdateBook(book entity.Book) (entity.Book, error) {
	book.UserID = 2
	if book, err := repository.UpdateBook(book); err == nil {
		return book, nil
	}
	return book, errors.New("book do not exists")
}

func DeleteBook(bookID uint64) error {
	if err := repository.DeleteBook(bookID); err == nil {
		return nil
	}
	return errors.New("book do not exists")
}
