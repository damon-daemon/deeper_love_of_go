package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		Title:  "Homage to Catalonia",
		Author: "George Orwell",
		Copies: 15,
	}
	exp := "Homage to Catalonia by George Orwell (copies: 15)"
	resp := books.BookToString(input)
	if exp != resp {
		t.Fatalf("expected %v but got %v", exp, resp)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	exp := []books.Book{
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
	resp := books.GetAllBooks()

	if !slices.Equal(exp, resp) {
		t.Fatalf("expected %#v but got %#v", exp, resp)
	}
}
