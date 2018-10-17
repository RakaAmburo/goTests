package main

import (
	"fmt"
	//"io"
	"net"
	"./pipe"
)

func main() {
	ln, err := net.Listen("tcp", ":4505")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println("new client")

	proxy, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		panic(err)
	}

	fmt.Println("proxy connected")
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	//io.Copy(src, dest)
	pipe.CopyBuffer(src, dest, nil)
}
