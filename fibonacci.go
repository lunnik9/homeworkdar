package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	a := 0
	b := 1
	for i := 0; i < 10; i++ {
		if a > b {
			mu.Lock()
			ch <- a
			Fork(ch)
			mu.Unlock()
			b += a
		} else {
			mu.Lock()
			ch <- b
			wg.Wait()
			Fork(ch)
			wg.Done()
			mu.Unlock()
			a += b
		}
	}
}
func Fork(ch chan int) {
	x := <-ch
	close(ch)
	ch1, ch2 := make(chan int), make(chan int)
	ch1 <- x
	ch2 <- x
	ToConsole(ch1)
	//	ToFile(ch2)
}

func ToConsole(ch chan int) {
	x := <-ch
	fmt.Println(x)
	close(ch)
}
func ToFile(ch chan int) {

}
