package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ch := make(chan int)
	ch1, ch2 := make(chan int), make(chan int)
	file, _ := os.Create("fibonacci.txt")
	a := 0
	b := 1
	for i := 0; i < 20; i++ {
		if a > b {
			go WriteToChannel(ch, a)
			Fork(ch, ch1, ch2, file)
			b += a
		} else {
			go WriteToChannel(ch, b)
			Fork(ch, ch1, ch2, file)
			a += b
		}
	}
}
func Fork(ch chan int, ch1 chan int, ch2 chan int, f *os.File) {
	x := <-ch
	go WriteToChannel(ch1, x)
	go WriteToChannel(ch2, x)
	ToConsole(ch1)
	ToFile(ch2, f)
}

func ToConsole(ch chan int) {
	x := <-ch
	fmt.Println(x)
}
func ToFile(ch chan int, f *os.File) {
	x := <-ch
	bs := []byte(strconv.Itoa(x))
	f.Write(bs)
	f.WriteString("\n")
}

func WriteToChannel(ch chan int, x int) {
	ch <- x
}
