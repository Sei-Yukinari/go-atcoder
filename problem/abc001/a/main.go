package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

	o := 0
	a := stdinInt()
	b := stdinInt()
	o = a - b
	stdout(o)

}

func stdinInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func stdout(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}
