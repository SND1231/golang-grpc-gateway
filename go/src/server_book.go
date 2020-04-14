package main

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	pb "test/book"
)

const (
	port = ":9002"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedBookServiceServer
}

// GET Book
func (s *server) GetBook(ctx context.Context, in *pb.GetBookIsbn) (*pb.GetBookStruct, error) {
	var book pb.Book
	book.Isbn = "isbntoha"
	book.Title = "This World"
	return &pb.GetBookStruct{Book: &book}, nil
}

// Create Book
func (s *server) CreateBook(ctx context.Context, in *pb.CreateBookStruct) (*pb.GetBookIsbn, error) {
	var book = *in
	fmt.Print(book.Isbn)
	fmt.Print(book.Title)
	return &pb.GetBookIsbn{Isbn: "2"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}