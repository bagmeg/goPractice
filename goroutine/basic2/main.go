package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	// Add 1 to wg, this will make wg to wait for 1 goroutine
	wg.Add(1)
	go func() {
		// notice wg not to wait
		defer wg.Done()
		fmt.Println("hi!! I'm goroutine")
	}()
	// wait until Done()
	wg.Wait()
	fmt.Println("hi!! I'm main routine")
}
