package main

import (
	"fmt" //"io"
	"net"

	"./pipe"
)

func main() {

	fmt.Println("Started.")
	conf := pipe.GetProps()
	fmt.Println(conf.SrcPort)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.SrcPort))
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
	conf := pipe.GetProps()

	proxy, err := net.Dial("tcp", fmt.Sprintf("%s:%d", conf.DestIp, conf.DestPort))
	if err != nil {
		panic(err)
	}

	fmt.Println("proxy connected")
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src net.Conn, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	//io.Copy(src, dest)
	pipe.CopyBuffer(src, dest, nil)
}
