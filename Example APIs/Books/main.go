package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       int32  `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int32  `json:"quantity"`
}

var books = []Book{
	{ID: 1, Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: 3, Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	{ID: 4, Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 3},
	{ID: 5, Title: "1984", Author: "George Orwell", Quantity: 4},
	{ID: 6, Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 7},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func addBooks(context *gin.Context) {
	var newBook Book

	err := context.BindJSON(&newBook)
	if err != nil {
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func searchBookByID(id int32) (*Book, error) {
	for index, Book := range books {
		if Book.ID == id {
			return &books[index], nil
		}
	}

	return nil, errors.New("Searched book not found!")
}

func getBookById(context *gin.Context) {
	id, parseError := strconv.ParseInt(context.Param("id"), 10, 32)
	book, searchError := searchBookByID(int32(id))

	if parseError != nil || searchError != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No book found with the given ID. Please enter a valid book ID"})
		return
	}

	context.IndentedJSON(http.StatusOK, book)
}

func issueBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter a valid book ID"})
		return
	}

	intID, parseError := strconv.ParseInt(id, 10, 32)
	book, searchError := searchBookByID(int32(intID))

	if parseError != nil || searchError != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found. Please enter a valid book ID"})
		return
	}

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not avaiable"})
		return
	}

	book.Quantity -= 1

	context.IndentedJSON(http.StatusOK, book)
}

func returnBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter a valid book ID"})
		return
	}

	intID, parseError := strconv.ParseInt(id, 10, 32)
	book, searchError := searchBookByID(int32(intID))

	if parseError != nil || searchError != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found. Please enter a valid book ID"})
		return
	}

	book.Quantity += 1
	context.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", addBooks)
	router.PATCH("/issue", issueBook)
	router.PATCH("/return", returnBook)

	router.Run("localhost:4000")
}
