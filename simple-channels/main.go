package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			// do something
		}
		pong <- fmt.Sprintf(strings.ToUpper(s))
	}
}

func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press enter (Q to quit)")

	for {
		// prompt the user
		fmt.Print("-> ")

		// get the user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput

		// wait for a response
		response := <-pong
		fmt.Println("Response:", response)
	}
	fmt.Println("Exiting.. closing channels!")
	close(ping)
	close(pong)
}
