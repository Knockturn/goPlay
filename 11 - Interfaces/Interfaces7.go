package main

import "fmt"

// start: https://youtu.be/YS4e4q9oBaU?t=19205
// Simplified version of previous examples
/*
If at least one of the methods is a pointer you have to use the &
*/

func main() {
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
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

type myWriterCloser struct{}

func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}
