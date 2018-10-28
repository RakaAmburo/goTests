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

func handleRequest4Serial(conn net.Conn) {
	fmt.Println("conecting serial")
	//conf := pipe.GetProps()

	proxy, err := pipe.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("proxy connected")
	go copyIO4Serial(conn, proxy)
	go copyIO4Serial(proxy, conn)

}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	//io.Copy(src, dest)
	pipe.CopyBuffer(src, dest, nil)
}

func copyIO4Serial(src, dest pipe.ReaderWriter) {
	defer src.Close()
	defer dest.Close()

	pipe.CopyBuffer(src, dest, nil)
}
