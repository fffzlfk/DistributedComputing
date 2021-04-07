package main

import (
	pb "bookmanagement/bookinfomgt"
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func Test(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not conneect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBookInfoMgtClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run("AddBook", func(t *testing.T) {
		book1 := &pb.Book{
			Id:   1,
			Name: "Rain",
		}
		r, err := c.AddBook(ctx, book1)
		if err != nil {
			t.Fatalf("could not add: %v\n", err)
		}
		if !r.GetOk() {
			t.Fatal("AddBook 测试错误")
		}
	})

	t.Run("QueryById", func(t *testing.T) {
		book, err := c.QueryById(ctx, &pb.QueryByIdRequest{Id: 1})
		if err != nil {
			t.Fatalf("could not query: %v\n", err)
		}

		if book.Name != "Rain" {
			log.Fatal("QueryById 测试错误")
		}
	})

	t.Run("QueryByName", func(t *testing.T) {
		book2 := &pb.Book{
			Id:   2,
			Name: "Rain",
		}
		r, err := c.AddBook(ctx, book2)
		if err != nil {
			t.Fatalf("could not add: %v\n", err)
		}
		if !r.Ok {
			t.Fatalf("AddBook failed")
		}

		books, err := c.QueryByName(ctx, &pb.QueryByNameRequest{Name: "Rain"})
		if err != nil {
			t.Fatalf("could not query: %v", err)
		}

		if len(books.Books) != 2 {
			t.Fatal("QueryByName 测试错误")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		rsp, err := c.Delete(ctx, &pb.DeleteRequest{Id: 1})
		if err != nil {
			log.Println("could not delete:", err)
		}

		if !rsp.Ok {
			t.Fatal("Delete 测试错误")
		}

		books, err := c.QueryByName(ctx, &pb.QueryByNameRequest{Name: "Rain"})
		if err != nil {
			t.Fatal("could not query:", err)
		}

		if len(books.Books) != 1 {
			t.Fatal("Delete 测试错误")
		}
	})
}
