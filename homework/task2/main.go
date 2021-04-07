package main

import (
	"bookmanagement/server"
	"flag"
)

var (
	port = flag.String("port", ":50051", "set port")
)

func main() {
	server.Start(*port)
}
