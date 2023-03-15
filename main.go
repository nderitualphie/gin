package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Chronicles of Narnia", Author: "prince caspian", Quantity: 2},
	{ID: "2", Title: "Spare", Author: "prince Harry", Quantity: 3},
	{ID: "3", Title: "Golden bells", Author: " unknown", Quantity: 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func getBookbyId(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func bookbyId(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookbyId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
func CheckOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing book"})
		return
	}
	book, err := getBookbyId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing id query parameter"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "book not available",
		})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing book"})
		return
	}
	book, err := getBookbyId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing id query parameter"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookbyId)
	router.PATCH("/checkout", CheckOutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:9003")
}
