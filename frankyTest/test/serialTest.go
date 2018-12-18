package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

var ventLightOn = []byte("D1TIT1&489F1039T1961H3936S#8HFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFZ|")
var ventLightOff = []byte("D1TIT1&500F1000T2000H3900S#8HFT3FT27FSFT3FT27FSFT3FT27FSFT3FT27FSFT3FT27FSFT3FT27FSFT3FT27FSFT3FT27FSFT3FT27FSFT3FT27FZ|")
var velador = []byte("D1TIT2&109L109L110L109L110L6252R#3522652532553525352532562L2522652622562535253522652L35226526225RR622652L2523552622562RR3562L26225R35RR622652L26135R35RR6225620|")
var luzCama = []byte("D1TIT2&108L108L108L109L109L109L109L26225625R#2522562523552625235526162L261355RR262LR23561626135616262LR23561626135625262L2523462622562616225626252LR23561626135626162L2613562523561626135626162LR235625262256262520|")
var mute = []byte("D1TIT1&838F1680T19012H#4FT12FT2F2TFH4FT12FT2F2TFS|")

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        "/dev/cu.usbmodem143241",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	go readFromArdu(port)

	time.Sleep(5 * time.Second)

	fmt.Println("ejecutando")
	// Write 4 bytes to the port.
	b := []byte("D1TIT1&489F1039T1961H3936S#8HFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFSFT3FT17FTFTFTFTFTFZ|")
	n, err := port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")

	time.Sleep(5 * time.Second)

}

func readFromArdu(port io.ReadCloser) {

	for {
		buf := make([]byte, 32)
		n, err := port.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from serial port: ", err)
			}
		} else {
			buf = buf[:n]
			fmt.Print(string(buf))
		}
	}

}
