package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sevendycom/poc-htmx-alpine/internal/api"
)

func main() {

	port := flag.String("port", ":8080", "Port to serve")
	flag.Parse()

	fmt.Println("Server running on port", *port)
	server := api.NewServer(*port)
	log.Fatal(server.Start())

}
