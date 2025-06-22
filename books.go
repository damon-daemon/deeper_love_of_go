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

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func GetAllBooks(catalog map[string]Book) []Book {
	return slices.Collect(maps.Values(catalog))
}

func GetBook(catalog map[string]Book, id string) (Book, bool) {
	book, ok := catalog[id]
	return book, ok
}

func AddBook(catalog map[string]Book, book Book) {
	catalog[book.ID] = book
}

func GetCatalog() map[string]Book {
	return map[string]Book{
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
}
