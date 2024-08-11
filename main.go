package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vaibhavp1964/go-redis/pkg"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	pkg.Run(port)
	log.Println("Starting Redis server on port:", port)
}
