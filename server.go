package main

import (
	"fmt"
	"net"

	"github.com/atotto/clipboard"
)

type BlockingServerBase struct {
	timeout int
	buffer  int
}

func NewBlockingServerBase(timeout, buffer int) *BlockingServerBase {
	return &BlockingServerBase{
		timeout: timeout,
		buffer:  buffer,
	}
}

func (bs *BlockingServerBase) Accept(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server started:", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go bs.handleConnection(conn)
	}
}

func (bs *BlockingServerBase) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, bs.buffer)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		message := string(buf[:n])
		bs.Respond(message)
	}

}

func (bs *BlockingServerBase) Respond(message string) {
	clipboard.WriteAll(message)
	fmt.Println(message)
	fmt.Println("received ->", message)
}

type InetServer struct {
	*BlockingServerBase
	host string
	port int
}

func NewInetServer(host string, port int) *InetServer {
	return &InetServer{
		BlockingServerBase: NewBlockingServerBase(500, 1024),
		host:               host,
		port:               port,
	}
}

func (is *InetServer) Start() {
	address := fmt.Sprintf("%s:%d", is.host, is.port)
	is.Accept(address)
}

func main() {
	is := NewInetServer("10.63.60.215", 8080)
	for {
		is.Start()
	}
}
