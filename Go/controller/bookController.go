package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasir16/simpe-api-rbac/structs"
)

// GetBooks is for get list book
func (idb *InDB) GetBooks(c *gin.Context) {
	var (
		books  []structs.Book
		result gin.H
	)

	idb.DB.Find(&users)
	if len(books) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": books,
			"count":  len(books),
		}
	}

	c.JSON(http.StatusOK, result)

}

// GetBook is get single Book
func (idb *InDB) GetBook(c *gin.Context) {
	var (
		book   structs.Book
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": book,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// CreateBook is creating Book
func (idb *InDB) CreateBook(c *gin.Context) {
	var (
		book   structs.Book
		result gin.H
	)
	err := c.Bind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
			"error":   err.Error(),
		})
	} else {
		idb.DB.Create(&book)
		result = gin.H{
			"result": book,
		}
		c.JSON(http.StatusOK, result)
	}

}

// UpdateBook is update single book by id
func (idb *InDB) UpdateBook(c *gin.Context) {
	var updateBook structs.Book
	errParam := c.Bind(&updateBook)
	if errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	id := c.Param("id")
	var (
		book    structs.Book
		newBook structs.Book
		result  gin.H
	)
	err := idb.DB.First(&book, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newBook.FullName = updateBook.FullName
	newBook.Username = updateBook.Username
	newBook.Phone = updateBook.Phone
	newBook.Address = updateBook.Address
	err = idb.DB.Model(&book).Updates(newBook).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"data":   newBook,
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// DeleteBook is deleting single book by id
func (idb *InDB) DeleteBook(c *gin.Context) {
	var (
		book   structs.Book
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&book, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&book).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
