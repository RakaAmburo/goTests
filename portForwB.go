package main

import (
	"fmt" //"io"
	"net"
	"time"

	"./pipe"
	"github.com/jacobsa/go-serial/serial"
)

func main() {

	fmt.Println("Started.")
	conf := pipe.GetProps()
	fmt.Println(conf.SrcPort)

	options := serial.OpenOptions{
		PortName:               conf.SerialPortB,
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
	ln, err := pipe.Open(options)
	if err != nil {
		panic(err)
	}

	p, er := net.Listen("tcp", fmt.Sprintf(":%d", 8818))
	if er != nil {
		panic(er)
	}

	go handleRequest2(ln)

	for {
		_, er2 := p.Accept()
		if er2 != nil {
			panic(er2)
		}
	}

}

func handleRequest2(conn *pipe.Usb) {
	fmt.Println("new client")
	conf := pipe.GetProps()

	for {

		proxy, err := net.Dial("tcp", fmt.Sprintf("%s:%d", conf.DestIp, conf.DestPort))
		if err != nil {
			panic(err)
		}
		conn.Listen(proxy)

		fmt.Println("proxy connected")
		go copyIO4Serial2(conn, proxy)
		time.Sleep(200 * time.Millisecond)
		go copyIO4Serial2(proxy, conn)

	}

}

func copyIO4Serial2(src, dest pipe.ReaderWriter) {
	//defer src.Close()
	//defer dest.Close()

	pipe.CopyBuffer(src, dest, nil)

}
