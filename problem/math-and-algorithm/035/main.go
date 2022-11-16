package main

import (
	"bufio"
	"fmt"
	"math"
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

type Pattern int

const (
	A Pattern = iota + 1 //一方の円が他方の円を完全に含み、2つの円は接していない
	B                    //一方の円が他方の円を完全に含み、2つの円は接している
	C                    //2つの円が互いに交差する
	D                    //2つの円の内部に共通部分は存在しないが、2つの円は接している
	E                    //2つの円の内部に共通部分は存在せず、2つの円は接していない
)

func main() {
	io := NewIo()
	defer io.Flush()

	c1 := io.NextInts(3)
	c2 := io.NextInts(3)

	var bigCircle []int
	var smallCircle []int
	if c1[2] >= c2[2] {
		bigCircle = c1
		smallCircle = c2
	} else {
		bigCircle = c2
		smallCircle = c1
	}
	x := float64(bigCircle[0]) - float64(smallCircle[0])
	y := float64(bigCircle[1]) - float64(smallCircle[1])
	d := math.Abs(math.Sqrt(x*x + y*y))

	r1 := float64(bigCircle[2])
	r2 := float64(smallCircle[2])

	if d > r1+r2 {
		io.Println(E)
	} else if d == r1+r2 {
		io.Println(D)
	} else if r1-r2 < d && d < r1+r2 {
		io.Println(C)
	} else if d == r1-r2 {
		io.Println(B)
	} else if d < r1-r2 {
		io.Println(A)
	}
}
