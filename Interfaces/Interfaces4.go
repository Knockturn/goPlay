package main

import (
	"bytes"
	"fmt"
	"io"
)

// Start: https://youtu.be/YS4e4q9oBaU?t=18715
// Same as Interfaces3 but there is a type conversion on line 18

func main() {
	fmt.Println("Interfaces3")
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()

	// Type conversion below From WriterCloser to BufferedWriterCloser
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	bwc2 := wc.(io.Reader) // Will fail because we dont implement the reader method.
	fmt.Println(bwc2)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// Progress: https://youtu.be/YS4e4q9oBaU?t=18818
