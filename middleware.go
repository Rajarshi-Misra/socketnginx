package main

import (
	"log"
	"net"
)

type MiddlewareFunc func([]byte) []byte

func loggingMiddleware(data []byte) []byte {
	log.Printf("Data. %s", string(data))
	return data
}

func applyMiddleware(data []byte, middlewares ...MiddlewareFunc) []byte {
	for _, m := range middlewares {
		data = m(data)
	}
	return data
}

func proxyWithMiddleware(src, dst net.Conn, middlewares ...MiddlewareFunc) {
	buf := make([]byte, 4096)
	for {
		n, err := src.Read(buf)

		if err != nil {
			break
		}

		data := buf[:n]
		data = applyMiddleware(data, middlewares...)

		_, err = dst.Write(data)

		if err != nil {
			break
		}
	}
}
