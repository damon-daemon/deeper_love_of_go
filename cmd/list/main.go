package main

import (
	"books"
	"fmt"
)

func main() {
	book := books.Book{
		Title:  "Homage to Catalonia",
		Author: "George Orwell",
		Copies: 15,
	}
	fmt.Println(books.BookToString(book))
}
