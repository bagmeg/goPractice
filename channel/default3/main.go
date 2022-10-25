package main

import (
	"fmt"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
		time.Sleep(time.Second)
	}

}

func merge(a, b <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					break
				}
				ch <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					break
				}
				ch <- v
			}
		}
	}()
	return ch
}

func asChan(val ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range val {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
