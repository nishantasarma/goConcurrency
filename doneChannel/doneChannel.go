package main

import (
	"fmt"
	"time"
)

// Done is a read ony channel
func doWork(done <-chan bool) {

	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work")
		}
	}

}

func main() {

	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)

}
