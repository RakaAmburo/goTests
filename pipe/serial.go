package pipe

import (
	"fmt"
	"io"
	"os"

	"github.com/jacobsa/go-serial/serial"
)

var rwc io.ReadWriteCloser

// Write to serial
func Write(bytes []byte) (n int, err error) {

	rwc.Write(bytes)

	return
}

// Close serial
func Close() {
	rwc.Close()
}

//Open the serial
func Open() {
	options := serial.OpenOptions{}
	rwc, err := serial.Open(options)
	if err != nil {
		fmt.Println("Error opening serial port: ", err)
		os.Exit(-1)
	} else {
		defer rwc.Close()
	}
}
