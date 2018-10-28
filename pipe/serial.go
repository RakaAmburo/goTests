package pipe

import (
	"io"

	"github.com/jacobsa/go-serial/serial"
)

//Usb serial
type Usb struct {
	rwc io.ReadWriteCloser
}

// Write to serial
func (s Usb) Write(bytes []byte) (n int, err error) {

	s.rwc.Write(bytes)

	return
}

// Read from serial
func (s Usb) Read(b []byte) (n int, err error) {
	return
}

// Close serial
func (s Usb) Close() (err error) {
	return s.rwc.Close()
}

//Open the serial
func Open() (s Usb, err error) {
	options := serial.OpenOptions{}
	rwc, err := serial.Open(options)
	s = Usb{rwc}

	return
}
