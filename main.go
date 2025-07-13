package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	ln, err := net.Listen("tcp", ":8000")
	checkError(err)
	fmt.Println("Listening on 8000")
	conn, err := ln.Accept()
	checkError(err)
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		checkError(err)
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}

func main() {
	go func() {
		fmt.Println("Starting server...")
		startServer()
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Starting client...")
	StartClient()
}
