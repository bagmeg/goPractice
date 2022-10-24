package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hi!! I'm goroutine")
	}()
	fmt.Println("hi!! I'm main routine")
	time.Sleep(time.Second)
}
