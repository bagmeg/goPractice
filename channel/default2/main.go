package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan bool)
	go f(c)

	go func() {
		for done := false; !done; {
			select {
			default:
				println("default")
			case <-ctx.Done():
				println("sent")
				done = true
			}
			time.Sleep(time.Second)
			c <- true
		}
	}()
	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 1)
}

func f(c <-chan bool) {
	for <-c {
		fmt.Println("got something")
	}
}
