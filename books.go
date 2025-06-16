package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

var catalog = []Book{
	{
		Title:  "1984",
		Author: "George Orwell",
		Copies: 10,
		ID:     "abc",
	},
	{
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
	return catalog
}

func GetBook(id string) (Book, bool) {
	for _, book := range catalog {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}
