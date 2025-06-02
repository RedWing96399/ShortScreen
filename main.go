package main

import (
	"log"

	"github.com/RedWing96399/shortscreen/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	server.Run(server.GetMux())
}
