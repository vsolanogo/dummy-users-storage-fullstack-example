package main

import (
	"log"
	"userproject/rest"
)

func StartServer() {
	rest.RunAPI("127.0.0.1:3000")
}

func main() {
	log.Println("Main log....")
	StartServer()
}
