package main

import (
	"bytes"
	"fmt"
)

// Start: https://youtu.be/YS4e4q9oBaU?t=18482

/*
Note:
Below is not the optimal way to do it but example is done in a Playground for Go.
And is setup so it could run in there.

Its a complicated example to show composing interfaces together
As long as you implement all of the methods on the embedded interfaces
then you actually implement the composed interface as well.
*/

func main() {
	fmt.Println("Interfaces3")
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()
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
