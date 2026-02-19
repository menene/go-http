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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	requestLine = strings.TrimSpace(requestLine)
	fmt.Println("Request:", requestLine)

	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		return
	}

	method := parts[0]
	path := parts[1]

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
	}

	routeRequest(conn, method, path)
}

func routeRequest(conn net.Conn, method, path string) {
	if method != "GET" {
		sendResponse(conn, "405 Method Not Allowed", "<h1>405 - Method Not Allowed</h1>")
		return
	}

	switch path {
		case "/":
			sendResponse(conn, "200 OK", "<h1>Home</h1><p>Welcome to branch 02</p>")
		case "/about":
			sendResponse(conn, "200 OK", "<h1>About</h1><p>This is manual routing.</p>")
		default:
			sendResponse(conn, "404 Not Found", "<h1>404 - Page Not Found</h1>")
	}
}

func sendResponse(conn net.Conn, status, body string) {
	response := fmt.Sprintf(
		"HTTP/1.1 %s\r\nContent-Type: text/html\r\nConnection: close\r\n\r\n%s",
		status,
		body,
	)

	conn.Write([]byte(response))
}
