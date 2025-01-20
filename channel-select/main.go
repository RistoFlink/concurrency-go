package main

import (
	"fmt"
	"time"
)

func serverOne(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "Message from serverOne"
	}
}
func serverTwo(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Message from serverTwo"
}
func main() {
	fmt.Println("Select w/ channels")
	fmt.Println("------------------")

	channelOne := make(chan string)
	channelTwo := make(chan string)

	go serverOne(channelOne)
	go serverTwo(channelTwo)

	for {
		select {
		// if multiple matches - select will just pick one randomly
		case s1 := <-channelOne:
			fmt.Println("Case one:", s1)
		case s2 := <-channelOne:
			fmt.Println("Case two:", s2)
		case s3 := <-channelTwo:
			fmt.Println("Case three:", s3)
		case s4 := <-channelTwo:
			fmt.Println("Case four:", s4)
			//default:
			// avoiding a deadlock
		}

	}
}
