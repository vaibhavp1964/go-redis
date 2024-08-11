package pkg

import (
	"fmt"
	"net"
)

const (
	RedisConnectionProtocol = "tcp"
)

func Run(port string) {
	connectionAddress := "localhost:" + port

	listener, err := net.Listen(RedisConnectionProtocol, connectionAddress)
	if err != nil {
		fmt.Println("error while setting up listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Accepting TCP connections on port:", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept incoming connection:", err)
			return
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error wbile reading incoming buffer:", err)
			return
		}

		fmt.Println("Incoming Buffer:", string(buffer[:n]))

		resp := []byte(ProcessCommand(string(buffer[:n])))

		_, err = conn.Write(resp)
		if err != nil {
			fmt.Println("error while sending data back to client")
		}

		return
	}
}

// ProcessCommand function  
func ProcessCommand(command string) string {
	return "Command registered"
}
