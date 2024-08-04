package main

import (
	"log"
)

func main() {
	server := NewAPIServer(":8080")
	log.Fatal(server.Start())
}
