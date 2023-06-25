package main

import (
	"fmt"
	"time"
)

func numbers(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("written on the channel : %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	// closing the channel in order to prevent deadlock
	close(n)
}

func main() {
	// second parameters represents the size of the channel
	channelNumber := make(chan int, 3)

	go numbers(channelNumber)

	for value := range channelNumber {
		fmt.Printf("read from the channel: %d\n", value)
		time.Sleep(time.Millisecond * 180)
	}

	fmt.Println("End of execution !")
}
