package main

import (
	pb "bookmanagement/bookinfomgt"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not conneect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBookInfoMgtClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	book1 := &pb.Book{
		Id:   1,
		Name: "Rain",
	}

	book2 := &pb.Book{
		Id:   2,
		Name: "Rain",
	}

	r, err := c.AddBook(ctx, book1)
	if err != nil {
		log.Printf("could not add: %v\n", err)
	}

	fmt.Println(r.GetOk())

	r, err = c.AddBook(ctx, book2)
	if err != nil {
		log.Printf("could not add: %v\n", err)
	}

	fmt.Println(r.GetOk())

	book, err := c.QueryById(ctx, &pb.QueryByIdRequest{Id: 1})
	if err != nil {
		log.Printf("could not query: %v\n", err)
	}

	fmt.Println(book.GetName())

	books, err := c.QueryByName(ctx, &pb.QueryByNameRequest{Name: "Rain"})
	if err != nil {
		log.Println("could not query:", err)
	}

	fmt.Println(books)

	rsp, err := c.Delete(ctx, &pb.DeleteRequest{Id: 1})
	if err != nil {
		log.Println("could not delete:", err)
	}

	fmt.Println(rsp.GetOk())

	books, err = c.QueryByName(ctx, &pb.QueryByNameRequest{Name: "Rain"})
	if err != nil {
		log.Println("could not query:", err)
	}

	fmt.Println(books)
}
