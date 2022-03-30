package main

import (
	"context"
	"fmt"

	pb "go.protobuf.foo.alis.exchange/foo/br/resources/books/v1"
)

func main() {
	// list books all books
	_ = listBooks()

	// get book with name "book/c4a2"
	_ = getBook("books/c4a2")
}

func listBooks() error {
	// List available books
	allBooks, err := booksClient.ListBooks(context.Background(), &pb.ListBooksRequest{})
	if err != nil {
		return fmt.Errorf("could not list books: %v", err)
	}

	for _, book := range allBooks.Books {
		fmt.Printf("%s\n", book.DisplayName)
	}

	return nil
}

func getBook(bookName string) error {
	req := pb.GetBookRequest{Name: bookName}

	book, err := booksClient.GetBook(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("could not get book: %v", err)
	}
	fmt.Printf("%s\n", book)

	return nil
}
