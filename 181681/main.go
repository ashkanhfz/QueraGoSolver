package library

import "strings"

type Library struct {
	capacity int
	books    map[string]string
}

func NewLibrary(capacity int) *Library {
	return &Library{
		capacity: capacity,
		books:    make(map[string]string),
	}
}

func (library *Library) AddBook(name string) string {
	bookName := strings.ToLower(name)

	if _, exists := library.books[bookName]; exists {
		return "The book is already in the library"
	}

	if len(library.books) >= library.capacity {
		return "Not enough capacity"
	}

	library.books[bookName] = ""
	return "OK"
}

func (library *Library) BorrowBook(bookName, personName string) string {
	book := strings.ToLower(bookName)

	borrower, exists := library.books[book]
	if !exists {
		return "The book is not defined in the library"
	}

	if borrower != "" {
		return "The book is already borrowed by " + borrower
	}

	library.books[book] = personName
	return "OK"
}

func (library *Library) ReturnBook(bookName string) string {
	book := strings.ToLower(bookName)

	borrower, exists := library.books[book]
	if !exists {
		return "The book is not defined in the library"
	}

	if borrower == "" {
		return "The book has not been borrowed"
	}

	library.books[book] = ""
	return "OK"
}
