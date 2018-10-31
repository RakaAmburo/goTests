package pipe

import (
	//"log"
	"errors"
	"fmt"
)

var ErrShortWrite = errors.New("short write")
var EOF = errors.New("EOF")

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() (err error)
}

type ReaderWriter interface {
	Reader
	Writer
	Closer
}

type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}

type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

type LimitedReader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func CopyBuffer(src Reader, dst Writer, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	/* if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	} */
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	/* if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	} */
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)

		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				fmt.Println(err)
				break
			}
			if nr != nw {
				err = ErrShortWrite
				fmt.Printf("%d:%d\n", nr, nw)
				break
			}
		}
		if er != nil {
			//fmt.Println(er)
			if er != EOF {
				err = er
				//fmt.Println(er)
				//fmt.Println(time.Now().Format(time.RFC850))
				//break
			} else {
				//fmt.Println(time.Now().Format(time.RFC850))
				//time.Sleep(200 * time.Millisecond)
			}

			break
		}
	}
	return written, err
}
