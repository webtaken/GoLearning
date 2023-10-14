package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	reader io.Reader
	limit  int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.limit {
		p = p[0:r.limit]
	}

	n, err = r.reader.Read(p)
	r.limit -= int64(n)
	return
}

func main() {
	l := LimitReader(strings.NewReader("1234567890"), 5)
	n, err := l.Read([]byte("123456789"))
	fmt.Println(n, err)
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
