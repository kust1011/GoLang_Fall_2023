package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID    int    `json:"id"`
    Name  string `json:"name"`
    Pages int    `json:"pages"`
}

var id int = 1

var bookshelf = []Book{
	Book{ID: 1, Name: "Blue Bird", Pages: 500},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, bookshelf)
	fmt.Println(bookshelf)
}

func getBook(c *gin.Context) {
	var bookID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

	for _, book := range bookshelf {
        if book.ID == bookID {
            c.JSON(http.StatusOK, book)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, b := range bookshelf {
		if b.Name == newBook.Name {
			c.JSON(http.StatusConflict, gin.H{"message": "duplicate book name"})
			return
		}
	}

	id++
	newBook.ID = id
	bookshelf = append(bookshelf, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	var bookID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for i, book := range bookshelf {
		if book.ID == bookID {
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func updateBook(c *gin.Context) {
	var bookID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for i, book := range bookshelf {
		if book.ID == bookID {
			var newBook Book
			if err := c.ShouldBindJSON(&newBook); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			for _, b := range bookshelf {
				if b.Name == newBook.Name && b.ID != bookID{
					c.JSON(http.StatusConflict, gin.H{"message": "duplicate book name"})
					return
				}
			}
			
			bookshelf[i].Name = newBook.Name
            bookshelf[i].Pages = newBook.Pages
			c.JSON(http.StatusOK, bookshelf[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true

	r.GET("/bookshelf",	getBooks)
	r.GET("/bookshelf/:id",	getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)

	err := r.Run(":8087")
	if err != nil {
		return
	}
}
