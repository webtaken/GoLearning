package main

import "io"

type CounterWriter struct {
	counter int64
	writer  io.Writer
}

// must be pointer type in order to count
func (cw *CounterWriter) Write(p []byte) (int, error) {
	cw.counter += int64(len(p))
	return cw.writer.Write(p)
}

// newWriter is a Writer Wrapper, return original Writer
// and a Counter which record bytes have written
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CounterWriter{0, w}
	return &cw, &cw.counter
}

func main() {

}
