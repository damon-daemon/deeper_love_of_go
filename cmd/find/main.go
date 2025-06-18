package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <book_id>")
		return
	}
	ID := os.Args[1]
	book, ok := books.GetBook(ID)
	if !ok {
		fmt.Println("Book not found")
		return
	}
	fmt.Println(books.BookToString(book))
}
