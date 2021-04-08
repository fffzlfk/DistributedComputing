package main

import (
	pb "bookmanagement/bookinfomgt"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type commandID int

const (
	CMD_ADD_BOOK commandID = iota
	CMD_QUERY_BY_ID
	CMD_QUERY_BY_NAME
	CMD_DELETE
)

type command struct {
	id   commandID
	args []string
}

const (
	address = "localhost:50051"
)

type client struct {
	pbc      *pb.BookInfoMgtClient
	commands chan command
}

func (c *client) ReadInput() {
	for {
		msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\n")
		args := strings.Split(msg, " ")
		cmd := strings.Trim(args[0], " ")
		switch cmd {
		case "add":
			if len(args) < 3 {
				fmt.Println("命令格式错误")
				continue
			}
			c.commands <- command{
				id:   CMD_ADD_BOOK,
				args: args,
			}
		case "qbi":
			c.commands <- command{
				id:   CMD_QUERY_BY_ID,
				args: args,
			}
		case "qbn":
			c.commands <- command{
				id:   CMD_QUERY_BY_NAME,
				args: args,
			}
		case "del":
			c.commands <- command{
				id:   CMD_DELETE,
				args: args,
			}
		default:
			fmt.Printf("unknown command: %s\n", cmd)
		}
	}

}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not conneect: %v", err)
	}
	defer conn.Close()

	pbc := pb.NewBookInfoMgtClient(conn)

	client := &client{
		pbc:      &pbc,
		commands: make(chan command),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go client.ReadInput()
	fmt.Println(
		`***书籍信息管理系统***
添加图书：	add [id] [name]
按id查询：	qbi [id]
按name查询：	qbn [name]
删除图书：	del [id]`)
	fmt.Println()
	for {
		cmd := <-client.commands
		switch cmd.id {
		case CMD_ADD_BOOK:
			id, _ := strconv.ParseInt(cmd.args[1], 10, 64)
			book := &pb.Book{
				Id:   id,
				Name: cmd.args[2],
			}
			r, err := (*(*client).pbc).AddBook(ctx, book)
			if err != nil {
				log.Printf("could not add: %v\n", err.Error())
			} else if r.Ok {
				fmt.Println("添加成功")
			}
		case CMD_QUERY_BY_ID:
			id, _ := strconv.ParseInt(cmd.args[1], 10, 64)
			book, err := (*(*client).pbc).QueryById(ctx, &pb.QueryByIdRequest{Id: id})
			if err != nil {
				log.Printf("could not query: %v\n", err)
			} else {
				fmt.Printf("the book is %v\n", book)
			}
		case CMD_QUERY_BY_NAME:
			name := cmd.args[1]
			books, err := (*(*client).pbc).QueryByName(ctx, &pb.QueryByNameRequest{Name: name})
			if err != nil {
				log.Println("could not query:", err)
			} else {
				fmt.Printf("查询到的%v书籍如下\n", name)
				for i, book := range books.Books {
					fmt.Printf("%dth: %v\n", i+1, book)
				}
			}
		case CMD_DELETE:
			id, _ := strconv.ParseInt(cmd.args[1], 10, 64)
			r, err := (*(*client).pbc).Delete(ctx, &pb.DeleteRequest{Id: id})
			if err != nil {
				log.Printf("could not delete: %v\n", err.Error())
			} else if r.Ok {
				fmt.Println("删除成功")
			}
		}
	}
}
