# Library Books System 
## This is a simple server system responsible for managing books in a library. 

## API Routes:
- GET `"/books"`: Returns the list of all the books in the library
- GET `"/books/:id`": Returns the information of the book with ID=`id`
- POST `"/books"`: Adds a new book to the list of all the books in the library
- PATCH `"/issue"`: Issues the book to someone. Reduces the quantity of the number of books in the library by 1
- PATCH `"/return"`: Returns the book to the library. Increments the quantity of the number of books in the library by 1

## Installation
- Clone this project and run `go run main.go` in your terminal to get started with the server.