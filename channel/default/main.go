package main

import "fmt"

func main() {
	c := make(chan bool, 1)

	for done := false; !done; {
		select {
		default:
			fmt.Println("done")
			done = true
		case c <- true:
			fmt.Println("sent")
		case <-c:
			fmt.Println("received")
			// this will block both c<-true and <-c
			// so default will be executed
			c = nil
		}
	}
	// result: sent, received, done
}
