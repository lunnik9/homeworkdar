package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	indexString := r.URL.Path
	indexString = strings.Replace(indexString, "/", "", -1)
	indexInt64, _ := strconv.ParseInt(indexString, 10, 32)
	index := int(indexInt64)
	bs, err := json.Marshal(FibonacciCounter(index))
	if err != nil {
		fmt.Println(err)
	}
	_, err = w.Write(bs)
	if err != nil {
		fmt.Println(err)
	}
}

func FibonacciCounter(index int) (f FibNumber) {
	a := 0
	b := 1
	for i := 0; i < index; i++ {
		if a > b {
			b += a
		} else {
			a += b
		}
	}
	if a > b {
		f.Current = a
		f.Prev = b
		f.Next = a + b
	} else {
		f.Current = b
		f.Prev = a
		f.Next = a + b
	}
	return f
}

type FibNumber struct {
	Current int
	Prev    int
	Next    int
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
