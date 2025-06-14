package books_test

import (
	"books"
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
