package main

import (
	"log"
	"userproject/rest"
)

func StartServer() {
	rest.RunAPI("0.0.0.0:3000")
}

func main() {
	log.Println("Main log....")
	StartServer()
}
