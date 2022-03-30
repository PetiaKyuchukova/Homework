package main

import (
	"io"
	"os"
)

func main() {
	ourReverseStringReader := NewReverseStringReader("Hello world!")
	io.Copy(os.Stdout, ourReverseStringReader)
}

type ReverseStringReader struct {
	data      []byte
	readIndex int64
}

func (r *ReverseStringReader) Read(p []byte) (n int, err error) {

	for i, j := 0, len(r.data)-1; i < j; i, j = i+1, j-1 {
		r.data[i], r.data[j] = r.data[j], r.data[i]
	}

	if r.readIndex >= int64(len(r.data)) {
		err = io.EOF
		return
	}

	n = copy(p, r.data[r.readIndex:])
	r.readIndex += int64(n)

	return
}

func NewReverseStringReader(input string) *ReverseStringReader {
	return &ReverseStringReader{data: []byte(input)}
}
