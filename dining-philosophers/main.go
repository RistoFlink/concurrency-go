package main

import (
	"fmt"
	"sync"
	"time"
)

// Philosopher is a struct which stores information about a philosopher
type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

// philosophers is a list of philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Kant", leftFork: 1, rightFork: 2},
	{name: "Hegel", leftFork: 2, rightFork: 3},
	{name: "Nietzsche", leftFork: 3, rightFork: 4},
}

// define some variables
var hunger = 3 // how many times a person eats
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// print out the welcome message
	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("-------------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()

	// print out the finished message
	fmt.Println("The table is empty again.")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex) // a pointer because you are never supposed to copy a mutex
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// start a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
}
