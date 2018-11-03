package main

import (
	"fmt" //"io"
	"net"
	"time"

	"./pipe"
	"github.com/jacobsa/go-serial/serial"
)

func main() {

	fmt.Println("Stargin usb to port pipe")
	conf := pipe.GetProps()

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

	usb, err := pipe.Open(options)
	if err != nil {
		panic(err)
	}

	port, err := net.Dial("tcp", fmt.Sprintf("%s:%d", conf.DestIp, conf.DestPort))
	if err != nil {
		panic(err)
	}

	usb.Listen(port)
	go copyIO4Serial2(usb, port)
	//time.Sleep(200 * time.Millisecond)
	go copyIO4Serial2(port, usb)

	time.Sleep(100000 * time.Millisecond)

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
