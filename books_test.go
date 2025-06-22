package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() books.Catalog {
	return books.Catalog{
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

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "Homage to Catalonia",
		Author: "George Orwell",
		Copies: 15,
		ID:     "xyz",
	}
	exp := "Homage to Catalonia by George Orwell (copies: 15)"
	resp := input.String()
	if exp != resp {
		t.Fatalf("expected %v but got %v", exp, resp)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
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
			ID:     "xyz",
		},
	}
	resp := catalog.GetAllBooks()
	slices.SortFunc(resp, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	slices.SortFunc(exp, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(resp, exp) {
		t.Fatalf("want %#v, got %#v", exp, resp)
	}
}

func TestGetBook_ReturnBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	exp := books.Book{
		ID:     "abc",
		Title:  "1984",
		Author: "George Orwell",
		Copies: 10,
	}
	resp, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatalf("book not found")
	}
	if exp != resp {
		t.Fatalf("expected %#v but got %#v", exp, resp)
	}
}

func TestGetBook_ReturnFalseWhenBookNotFoun(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("does-not-exist")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AppendsBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("book arleady present")
	}
	newBook := books.Book{
		ID:     "123",
		Title:  "Animal Farm",
		Author: "George Orwell",
		Copies: 8,
	}
	catalog.AddBook(newBook)

	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}
