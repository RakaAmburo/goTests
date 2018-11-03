package pipe

import (
	"fmt"
	"io"
	"math"
	"net"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

//Usb serial
type Usb struct {
	rwc io.ReadWriteCloser
}

// Write to serial
func (s Usb) Write(bytes []byte) (n int, err error) {

	pack := 32
	total := len(bytes)
	//fmt.Printf("total to be writen %d\n", total)
	nTotal := 0
	loops := math.Ceil(float64(total) / float64(pack))
	sent := 0
	for i := 0; i < int(loops); i++ {
		sent += pack
		if sent > total {
			rest := pack - (sent - total)
			sent -= pack
			sent += rest
			pack = rest
		}
		base := sent - pack
		top := sent

		nTotal, err = s.rwc.Write(bytes[base:top])
		n += nTotal
		//fmt.Println("aqui")
		//fmt.Println(top)
		//fmt.Println(total)
		time.Sleep(8000 * time.Microsecond) //3600
	}
	fmt.Printf("Write %d\n", n)
	return
}

//Listen to incoming bytes
func (s Usb) Listen(conn net.Conn) (n int, err error) {

	basePacksise := 32
	basePack := make([]byte, basePacksise)

	for {
		n, err := s.rwc.Read(basePack)

		if err != nil {
			//fmt.Println(err)
			if err != io.EOF {
				//fmt.Println(err)
				break
			} else {
				//fmt.Println("esperamos...")
				time.Sleep(200 * time.Millisecond)
			}
		}

		if n > 0 {
			n, err = conn.Write(basePack)
			break
		}
	}

	return n, err
}

// Read from serial
func (s Usb) Read(b []byte) (n int, err error) {

	basePacksise := 32
	basePack := make([]byte, basePacksise)
	count := 0

	for i := 1; i <= 320; i++ {

		n, err = s.rwc.Read(basePack)
		//fmt.Println(err)
		if err != nil {
			break
		} else {
			copy(b[count:], basePack[:n])
		}

		count += n
		time.Sleep(8000 * time.Microsecond)
	}
	fmt.Printf("Read %d\n", count+32)
	return count, err
}

// Close serial
func (s Usb) Close() (err error) {
	return s.rwc.Close()
}

//Open the serial
func Open(op serial.OpenOptions) (s *Usb, err error) {

	rwc, err := serial.Open(op)
	s = &Usb{rwc}

	return
}
