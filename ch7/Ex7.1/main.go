package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*w = WordCounter(count)
	return count, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*l = LineCounter(count)
	return count, nil
}

func main() {
	var c WordCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "1"
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "2"

	var d LineCounter
	d.Write([]byte("hello"))
	fmt.Println(d) // "1"
	d = 0          // reset the counter
	name = "\nDolly"
	fmt.Fprintf(&d, "hello, %show \nare you?", name)
	fmt.Println(d) // "3"
}
