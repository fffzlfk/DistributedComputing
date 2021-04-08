package main

import (
	pb "bookmanagement/bookinfomgt"
	"context"
	"errors"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
)

type BookInfoMgtServer struct {
	pb.UnimplementedBookInfoMgtServer
	books map[int64]*pb.Book
}

func NewBookInfoMgtServer(books map[int64]*pb.Book) *BookInfoMgtServer {
	return &BookInfoMgtServer{books: books}
}

func (s *BookInfoMgtServer) AddBook(ctx context.Context, book *pb.Book) (*pb.AddBookResponse, error) {
	if s.books[book.Id] != nil {
		return &pb.AddBookResponse{Ok: false}, errors.New("已经存在")
	}
	s.books[book.Id] = book
	log.Printf("added a book: id: %v, name: %v\n", book.Id, book.Name)
	return &pb.AddBookResponse{Ok: true}, nil
}

func (s *BookInfoMgtServer) QueryById(ctx context.Context, request *pb.QueryByIdRequest) (*pb.Book, error) {
	log.Println("querybyid:", request.Id)

	for k, v := range s.books {
		if k == request.Id {
			return v, nil
		}
	}
	return &pb.Book{Id: -1}, errors.New("没有找到")
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
	log.Println(request.Id, "deleted from books")

	for k, v := range s.books {
		if v.Id == request.Id {
			delete(s.books, k)
			return &pb.DeleteResponse{Ok: true}, nil
		}
	}
	return &pb.DeleteResponse{Ok: false}, errors.New("没有此书")
}

func (s *BookInfoMgtServer) ShowBooks(ctx context.Context, request *pb.ShowBooksRequest) (*pb.ShowBooksResponse, error) {
	books := make([]*pb.Book, 0)
	for _, v := range s.books {
		books = append(books, v)
	}
	return &pb.ShowBooksResponse{Books: books}, nil
}

var (
	port = flag.String("port", ":50052", "set port")
)

func init() {
	flag.Parse()
}

func main() {
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	log.Println("listen at", *port)

	s := grpc.NewServer()
	books := make(map[int64]*pb.Book)
	pb.RegisterBookInfoMgtServer(s, NewBookInfoMgtServer(books))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
