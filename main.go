package main

import (
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
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:9003")
}
