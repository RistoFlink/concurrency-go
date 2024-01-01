package main

import (
	"fmt"
	"time"
)

// Technically every program in Go has a GoRoutine since the main-function is itself one
// GoRoutines are very lightweight threads
func main() {
	go printSomething("This is the first thing to be printed!")

	time.Sleep(1 * time.Second)

	printSomething("This is the second thing to be printed!")
}

func printSomething(s string) {
	fmt.Println(s)
}
