package main

import (
	pb "bookmanagement/bookinfomgt"
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.String("port", ":50051", "set port")
)

type BookInfoMgtServer struct {
	pb.UnimplementedBookInfoMgtServer
	books map[int64]*pb.Book
}

func (s *BookInfoMgtServer) AddBook(ctx context.Context, book *pb.Book) (*pb.AddBookResponse, error) {
	s.books[book.Id] = book
	return &pb.AddBookResponse{Ok: true}, nil
}

func (s *BookInfoMgtServer) QueryById(ctx context.Context, request *pb.QueryByIdRequest) (*pb.Book, error) {
	for k, v := range s.books {
		if k == request.Id {
			return v, nil
		}
	}
	return &pb.Book{Id: -1}, ctx.Err()
}

func (s *BookInfoMgtServer) QueryByName(ctx context.Context, request *pb.QueryByNameRequest) (*pb.BookList, error) {
	books := make([]*pb.Book, 0)

	for _, v := range s.books {
		if v.Name == request.Name {
			books = append(books, v)
		}
	}

	return &pb.BookList{Books: books}, nil
}

func (s *BookInfoMgtServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	for k, v := range s.books {
		if v.Id == request.Id {
			delete(s.books, k)
			return &pb.DeleteResponse{Ok: true}, nil
		}
	}
	return &pb.DeleteResponse{Ok: false}, ctx.Err()
}

func init() {
	flag.Parse()
}

func main() {
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	books := make(map[int64]*pb.Book)
	pb.RegisterBookInfoMgtServer(s, &BookInfoMgtServer{books: books})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
