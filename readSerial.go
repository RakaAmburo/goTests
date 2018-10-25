package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jacobsa/go-serial/serial"
)

func usage() {
	fmt.Println("go-serial-test usage:")
	flag.PrintDefaults()
	os.Exit(-1)
}

func main() {
	fmt.Println("Go serial test")
	port := flag.String("port", "", "serial port to test (/dev/ttyUSB0, etc)")
	baud := flag.Uint("baud", 115200, "Baud rate")
	//txData := flag.String("txdata", "", "data to send in hex format (01ab238b)")
	even := flag.Bool("even", false, "enable even parity")
	odd := flag.Bool("odd", false, "enable odd parity")
	rs485 := flag.Bool("rs485", false, "enable RS485 RTS for direction control")
	rs485HighDuringSend := flag.Bool("rs485_high_during_send", false, "RTS signal should be high during send")
	rs485HighAfterSend := flag.Bool("rs485_high_after_send", false, "RTS signal should be high after send")
	stopbits := flag.Uint("stopbits", 1, "Stop bits")
	databits := flag.Uint("databits", 8, "Data bits")
	chartimeout := flag.Uint("chartimeout", 100, "Inter Character timeout (ms)")
	minread := flag.Uint("minread", 0, "Minimum read count")
	rx := flag.Bool("rx", false, "Read data received")

	flag.Parse()

	if *port == "" {
		fmt.Println("Must specify port")
		usage()
	}

	if *even && *odd {
		fmt.Println("can't specify both even and odd parity")
		usage()
	}

	parity := serial.PARITY_NONE

	if *even {
		parity = serial.PARITY_EVEN
	} else if *odd {
		parity = serial.PARITY_ODD
	}

	options := serial.OpenOptions{
		PortName:               *port,
		BaudRate:               *baud,
		DataBits:               *databits,
		StopBits:               *stopbits,
		MinimumReadSize:        *minread,
		InterCharacterTimeout:  *chartimeout,
		ParityMode:             parity,
		Rs485Enable:            *rs485,
		Rs485RtsHighDuringSend: *rs485HighDuringSend,
		Rs485RtsHighAfterSend:  *rs485HighAfterSend,
	}

	f, err := serial.Open(options)

	if err != nil {
		fmt.Println("Error opening serial port: ", err)
		os.Exit(-1)
	} else {
		defer f.Close()
	}

	if *rx {
		counter := 0
		for {
			buf := make([]byte, 1)
			n, err := f.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Println("Error reading from serial port: ", err)
				}
			} else {
				bait := buf[:n]
				if int(bait[0]) == 1 {
					fmt.Println("Rx: ", counter)
					counter = 0
				} else {
					counter++
				}

			}
		}
	}
}
