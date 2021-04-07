package server

import (
	pb "bookmanagement/bookinfomgt"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type BookInfoMgtServer struct {
	pb.UnimplementedBookInfoMgtServer
	books map[int64]*pb.Book
}

func (s *BookInfoMgtServer) AddBook(ctx context.Context, book *pb.Book) (*pb.AddBookResponse, error) {
	s.books[book.Id] = book
	log.Printf("added a book: name: %v, id: %v\n", book.Id, book.Name)
	return &pb.AddBookResponse{Ok: true}, nil
}

func (s *BookInfoMgtServer) QueryById(ctx context.Context, request *pb.QueryByIdRequest) (*pb.Book, error) {
	log.Println("querybyid:", request.Id)

	for k, v := range s.books {
		if k == request.Id {
			return v, nil
		}
	}
	return &pb.Book{Id: -1}, ctx.Err()
}

func (s *BookInfoMgtServer) QueryByName(ctx context.Context, request *pb.QueryByNameRequest) (*pb.BookList, error) {
	books := make([]*pb.Book, 0)

	log.Println("querybyname:", request.Name)

	for _, v := range s.books {
		if v.Name == request.Name {
			books = append(books, v)
		}
	}

	return &pb.BookList{Books: books}, nil
}

func (s *BookInfoMgtServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	fmt.Println(request.Id, "deleted from books")

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

func Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("listen at", port)

	s := grpc.NewServer()
	books := make(map[int64]*pb.Book)
	pb.RegisterBookInfoMgtServer(s, &BookInfoMgtServer{books: books})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
