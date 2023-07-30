package service

import (
	"bookapi/dto"
	"bookapi/entity"
	"bookapi/repository"
	"errors"
	"log"

	"github.com/mashingan/smapping"
)

func GetAllBooks() []entity.Book {
	return repository.GetAllBooks()
}

func InsertBook(bookDTO dto.BookCreatedDTO, userID uint64) dto.BookResponseDTO {
	book := entity.Book{}
	bookResponse := dto.BookResponseDTO{}

	err := smapping.FillStruct(&book, smapping.MapFields(&bookDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return bookResponse
	}

	book.UserID = userID
	book = repository.InsertBook(book)

	err = smapping.FillStruct(&bookResponse, smapping.MapFields(&book))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return bookResponse
	}

	return bookResponse
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

func IsAllowedToEdit(userID uint64, bookID uint64) bool {
	b := repository.GetTheBookUsingID(bookID)
	return userID == b.UserID
}
