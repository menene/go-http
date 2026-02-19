package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(err)
	}

	fmt.Println("Waiting for client on :80...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Client connected.")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		fmt.Println("Received:", line)
	}

	response := `HTTP/1.1 200 OK
Content-Type: text/html
Connection: close

<!DOCTYPE html>
<html>
<head>
    <title>Raw Go Server</title>
</head>
<body>
    <h1>Hello World</h1>
</body>
</html>`

	conn.Write([]byte(response))
}
