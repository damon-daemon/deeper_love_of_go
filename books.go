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

type Catalog map[string]Book

func (book Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func (catalog Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func (catalog Catalog) GetBook(id string) (Book, bool) {
	book, ok := catalog[id]
	return book, ok
}

func (catalog Catalog) AddBook(book Book) {
	catalog[book.ID] = book
}

func GetCatalog() Catalog {
	return Catalog{
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
