package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func NewIo() *Io {
	return &Io{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt()
	}
	return a
}

func (io *Io) NextFloat() float64 {
	i, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) Println(a ...interface{}) {
	_, err := fmt.Fprintln(io.writer, a...)
	if err != nil {
		return
	}
}

func (io *Io) Printf(format string, a ...interface{}) {
	_, err := fmt.Fprintf(io.writer, format, a...)
	if err != nil {
		return
	}
}

func (io *Io) PrintIntLn(a []int) {
	var b []interface{}
	for _, x := range a {
		b = append(b, x)
	}
	io.Println(b...)
}

func (io *Io) PrintStringLn(a []string) {
	var b []interface{}
	for _, x := range a {
		b = append(b, x)
	}
	io.Println(b...)
}

func Log(name string, value interface{}) {
	_, err := fmt.Fprintf(os.Stderr, "%s=%+v\n", name, value)
	if err != nil {
		return
	}
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	a := io.NextInts(n)

	io.PrintIntLn(mergeSort(a))
}

func mergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	mid := int(len(s) / 2)
	r := mergeSort(s[:mid])
	l := mergeSort(s[mid:])

	i, j := 0, 0
	var result []int
	for i < len(r) && j < len(l) {
		if r[i] > l[j] {
			result = append(result, l[j])
			j++
		} else {
			result = append(result, r[i])
			i++
		}
	}
	result = append(result, r[i:]...)
	result = append(result, l[j:]...)
	return result
}
