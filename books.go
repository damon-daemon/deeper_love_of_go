package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

var catalog = map[string]Book{
	"abc": {
		Title:  "1984",
		Author: "George Orwell",
		Copies: 10,
		ID:     "abc",
	},
	"xyz": {
		Title:  "Brave New World",
		Author: "Aldous Huxley",
		Copies: 5,
		ID:     "xyz",
	},
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func GetBook(id string) (Book, bool) {
	book, ok := catalog[id]
	return book, ok
}
