package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	Id    int     `binding:"required"`
	Title string  `binding:"required"`
	Price float64 `binding:"required"`
}

var books = []Book{
	{1, "Awesome Go", 100},
	{2, "Learn Go!", 90},
	{3, "Go Tricks", 60},
	{4, "The Power Of Go", 120},
	{5, "Go Concurrency", 110},
}

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/book/:id", getBookById)
	r.GET("/book", getAllBooks)
	r.POST("/book", addBook)
	return r
}

func main() {
	router := setUpRouter()

	_ = router.Run("localhost:8080")
}

func getAllBooks(ctx *gin.Context) {
	ctx.JSON(200, books)
}

func getBookById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	for _, book := range books {
		if book.Id == id {
			ctx.JSON(200, book)
			return
		}
	}

	ctx.JSON(404, gin.H{"message": fmt.Sprintf("book with id %d not found", id)})
}

func addBook(ctx *gin.Context) {
	var book Book
	err := ctx.BindJSON(&book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	books = append(books, book)

	ctx.JSON(http.StatusCreated, book)
}
