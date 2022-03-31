package main

import (
	"context"
	"fmt"

	// Import Books protobuf
	pb "go.protobuf.foo.alis.exchange/foo/br/resources/books/v1"
)

func main() {

	// List all books available
	_ = listBooks()

	// Get book with name "book/c4a2"
	_ = getBook("books/c4a2")
}

// listBooks function requests all available books
func listBooks() error {
	// Create a ListBooksRequest as defined in proto
	req := pb.ListBooksRequest{}

	// Call ListBooks method from booksClient
	allBooks, err := booksClient.ListBooks(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("could not list books: %v", err)
	}

	for _, book := range allBooks.Books {
		fmt.Printf("%s\n", book.DisplayName)
	}

	return nil
}

// Create a function that takes a bookName variable as a string input parameter.
// Create a getBookRequest as defined by the proto
// Use the getBookRequest as input parameter to the GetBook method on the booksClient
// Print the response

// getBook requests a book with bookName as input
func getBook(bookName string) error {
	// Create a GetBookRequest as defined in proto
	req := pb.GetBookRequest{Name: bookName}

	// Call GetBook method from booksClient with req as input
	book, err := booksClient.GetBook(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("could not get book: %v", err)
	}
	fmt.Printf("%s\n", book)

	return nil
}
