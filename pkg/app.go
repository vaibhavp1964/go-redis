package pkg

import (
	"fmt"
	"os"

  "github.com/joho/godotenv"
)

func Run() {
  godotenv.Load()

  port := os.Getenv("port")
  fmt.Println("Port:", port)
}
