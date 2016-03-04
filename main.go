package main

import (
	"log"

	"github.com/abhishekshivanna/pictomania/web"
)

func main() {
	web.StartServer()
	log.Println("Server started successfully")
}
