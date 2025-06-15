package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

var catalog = []Book{
	{
		Title:  "1984",
		Author: "George Orwell",
		Copies: 10,
	},
	{
		Title:  "Brave New World",
		Author: "Aldous Huxley",
		Copies: 5,
	},
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func GetAllBooks() []Book {
	return catalog
}
