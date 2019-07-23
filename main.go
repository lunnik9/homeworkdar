package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type Map map[string]int

func main() {
	b, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	m := make(Map)

	newLine := regexp.MustCompile("\n")
	matches := newLine.FindAllStringIndex(str, -1)
	nOfLines := len(matches)
	ch := make(chan *Map)
	line := ""
	for _, c := range str {
		if c == 10 || c == 13 {
			go m.countRunes(line, ch)
			line = ""
		} else {
			line += string(rune(c))
		}

	}

	for i := 0; i < nOfLines; i++ {
		b := <-ch
		fmt.Println(b)
	}
	fmt.Println(m)
}

func (m *Map) countRunes(s string, ch chan *Map) {
	for _, c := range s {
		if val, ok := (*m)[string(rune(c))]; ok {
			val++
		} else {
			(*m)[string(rune(c))] = 1
		}
	}
	ch <- m
}
