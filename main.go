package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vaibhavp1964/go-redis/pkg"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	pkg.Run(port)
}
