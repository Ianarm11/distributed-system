package main

import (
	server2 "github.com/Ianarm11/distributed-system/server"
	"log"
)

func main() {
	server := server2.NewHTTPServer(":8080")
	log.Fatalln(server.ListenAndServe())
}