package main

import (
	pb "bookmanagement/bookinfomgt"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net"
	"os"

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
	storeBooks(s.books)
	log.Printf("added a book: id: %v, name: %v\n", book.Id, book.Name)
	return &pb.AddBookResponse{Ok: true}, nil
}

func (s *BookInfoMgtServer) QueryById(ctx context.Context, request *pb.QueryByIdRequest) (*pb.Book, error) {
	log.Println("querybyid:", request.Id)

	book, exist := s.books[request.Id]
	if !exist {
		return nil, errors.New("没有找到")
	}

	return book, nil
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
			storeBooks(s.books)
			return &pb.DeleteResponse{Ok: true}, nil
		}
	}

	return &pb.DeleteResponse{Ok: false}, errors.New("没有此书")
}

func (s *BookInfoMgtServer) ShowBooks(ctx context.Context, request *pb.ShowBooksRequest) (*pb.ShowBooksResponse, error) {
	log.Println("show all books")
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

const dataFile = "server/data.json"

// 从文件恢复数据
func retrieveBooks() (map[int64]*pb.Book, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	books := make(map[int64]*pb.Book)

	var arr []*pb.Book
	err = json.NewDecoder(file).Decode(&arr)

	for _, v := range arr {
		books[v.Id] = v
	}

	return books, err
}

// 写回文件
func storeBooks(books map[int64]*pb.Book) error {
	var arr []*pb.Book
	for _, v := range books {
		arr = append(arr, v)
	}
	file, err := os.OpenFile(dataFile, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")
	err = enc.Encode(&arr)
	return err
}

func main() {
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	defer lis.Close()
	log.Println("listen at", *port)

	s := grpc.NewServer()
	books, err := retrieveBooks()
	if err != nil {
		log.Fatalf("failed to load data: %v\n", err)
	}

	pb.RegisterBookInfoMgtServer(s, NewBookInfoMgtServer(books))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
