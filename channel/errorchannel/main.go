// passing error to goroutine

package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

func main() {
	// create a context to control goroutine
	ctx, cancel := context.WithCancel(context.Background())

	// error channel to receive error
	errChan := someErrorFunc(ctx)

	// goroutine to write errors to channel
	go printError(ctx, errChan)

	// run program(goroutine) for 37 seconds
	time.Sleep(time.Second * 37)
	// cancel context to quit goroutine
	cancel()
	// wait for goroutine to quit - to see the logs
	time.Sleep(time.Second * 2)
}

// printError receives error from channel, and print it
func printError(ctx context.Context, errChan <-chan error) {
	for {
		select {
		// if receive error, print it
		case err := <-errChan:
			if err != nil {
				log.Println(err)
			} else {
				log.Println("received nil")
			}
		// if context is done, quit
		case <-ctx.Done():
			log.Println("Quit printError...")
			return
		}
	}
}

// someErrorFunc writes error or nil every 5 second  to channel
func someErrorFunc(ctx context.Context) <-chan error {
	// create a channel to send error
	errChan := make(chan error)
	// create a wait group to make sure goroutine starts
	wg := sync.WaitGroup{}

	// add 1 to wait group
	wg.Add(1)
	go func() {
		wg.Done()

		// timer to tick every 5 seconds
		ticker := time.NewTicker(time.Second * 5)
		// counter to count number of times
		ct := 0
		for {
			select {
			// every 5 seconds
			case <-ticker.C:
				if ct%2 == 0 {
					errChan <- errors.New("some error")
				} else {
					errChan <- nil
				}
				ct++
			// if context is done, quit
			case <-ctx.Done():
				log.Println("Quit someErrorFunc...")
				return
			}
		}
	}()
	// wait for goroutine to start
	wg.Wait()
	return errChan
}
