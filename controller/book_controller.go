package controller

import (
	"bookapi/dto"
	"bookapi/entity"
	"bookapi/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "select books",
		"books":   service.GetAllBooks(),
	})
}

func InsertBook(c *gin.Context) {
	var book dto.BookCreatedDTO
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	b := service.InsertBook(book, userID)
	c.JSON(200, gin.H{
		"message": "insert book",
		"book":    b,
	})
}

func GetBook(c *gin.Context) {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	book, err := service.GetBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "select book",
		"book":    book,
	})
}

func UpdateBook(c *gin.Context) {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var book entity.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	book.ID = bookID
	book, err = service.UpdateBook(book)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "update book",
		"book":    book,
	})
}

func DeleteBook(c *gin.Context) {
	bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if !service.IsAllowedToEdit(userID, bookID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this book",
		})
		return
	}

	err := service.DeleteBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "book deleted",
	})
}
