package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	go echo1(ch)
	go echo2(ch)
	go echo3(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Printf("%.15fs elapsed\n", time.Since(start).Seconds())
}

func echo1(ch chan<- string) {

	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Println(s)
	ch <- fmt.Sprintf("echo1 -> %.30fs", secs)

}

func echo2(ch chan<- string) {

	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("echo2 -> %.30fs", secs)

}

func echo3(ch chan<- string) {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("echo3 -> %.30fs", secs)
}
