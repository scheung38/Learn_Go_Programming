package main

import "fmt"

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Stopped")
			default:
				// …
			}
		}
	}()
	// …
	quit <- true
}
