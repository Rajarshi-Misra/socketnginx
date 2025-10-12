package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	listenAddr := flag.String("listen", ":8080", "Address to listen on")
	backendAddr := flag.String("backend", "localhost:9001", "Backend address to proxy to")
	flag.Parse()

	ln, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Proxy listening on %s, forwarding to %s\n", *listenAddr, *backendAddr)

	for {
		clientConn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go handleConnection(clientConn, *backendAddr)
	}
}

func handleConnection(clientConn net.Conn, backendAddr string) {
	backendConn, err := net.Dial("tcp", backendAddr)
	if err != nil {
		log.Println("Backend connection error:", err)
		clientConn.Close()
		return
	}
	defer clientConn.Close()
	defer backendConn.Close()

	// Bidirectional copy
	go proxyWithMiddleware(clientConn, backendConn, loggingMiddleware)
	proxyWithMiddleware(backendConn, clientConn, loggingMiddleware)
}
