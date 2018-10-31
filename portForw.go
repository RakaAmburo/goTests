package main

import (
	"fmt" //"io"
	"net"

	"./pipe"
	"github.com/jacobsa/go-serial/serial"
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

		go handleRequest4Serial(conn)
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

var usb *pipe.Usb

func handleRequest4Serial(conn net.Conn) {
	fmt.Println("conecting serial")

	if usb == nil {
		fmt.Println("creating usb serial")
		conf := pipe.GetProps()
		options := serial.OpenOptions{
			PortName:               conf.SerialPortA,
			BaudRate:               uint(115200),
			DataBits:               uint(8),
			StopBits:               uint(1),
			MinimumReadSize:        uint(0),
			InterCharacterTimeout:  uint(100),
			ParityMode:             serial.PARITY_NONE,
			Rs485Enable:            false,
			Rs485RtsHighDuringSend: false,
			Rs485RtsHighAfterSend:  false,
		}

		var err error

		usb, err = pipe.Open(options)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("usb connected")
	go copyIO4Serial(conn, usb)

	usb.Listen(conn)

	go copyIO4Serial(usb, conn)

}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	//io.Copy(src, dest)
	pipe.CopyBuffer(src, dest, nil)
}

func copyIO4Serial(src, dest pipe.ReaderWriter) {
	//defer src.Close()
	//defer dest.Close()

	pipe.CopyBuffer(src, dest, nil)
}
