package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "Homage to Catalonia",
		Author: "George Orwell",
		Copies: 15,
		ID:     "xyz",
	}
	exp := "Homage to Catalonia by George Orwell (copies: 15)"
	resp := books.BookToString(input)
	if exp != resp {
		t.Fatalf("expected %v but got %v", exp, resp)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	exp := []books.Book{
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
			ID:     "def",
		},
	}
	resp := books.GetAllBooks()

	if !slices.Equal(exp, resp) {
		t.Fatalf("expected %#v but got %#v", exp, resp)
	}
}

func TestGetBook_ReturnBook(t *testing.T) {
	t.Parallel()
	exp := books.Book{
		ID:     "abc",
		Title:  "1984",
		Author: "George Orwell",
		Copies: 10,
	}
	resp, ok := books.GetBook("abc")
	if !ok {
		t.Fatalf("book not found")
	}
	if exp != resp {
		t.Fatalf("expected %#v but got %#v", exp, resp)
	}
}

func TestGetBook_ReturnFalseWhenBookNotFoun(t *testing.T) {
	t.Parallel()
	_, ok := books.GetBook("does-not-exist")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}
