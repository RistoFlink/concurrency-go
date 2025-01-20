package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Preparing pizza #%d. It will take %d seconds\n", pizzaNumber, delay)
		// delay for a bit..
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** We accidentally burned the #%d\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza #%d is ready!\n", pizzaNumber)
		}
		p := &PizzaOrder{pizzaNumber, msg, success}
		return p
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of what pizza is being made
	var i = 0
	// run forever or until we receive a quit notification

	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// seed the RNG - not needed since go 1.20?
	// print the start message
	color.Green("Risto's pizzeria is open for business")
	color.Cyan("--------------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run the consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer ain't happy!")
			}
		} else {
			color.Cyan("Closed up shop for the day!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel", err)
			}
		}
	}
	// print the end message
	color.Cyan("----------------------")
	color.Cyan("Oh boy, quittin' time!")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attemps in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was not a good day..")
	case pizzasFailed > 6:
		color.Red("Could've been better..")
	case pizzasFailed >= 4:
		color.Yellow("It was okay..")
	case pizzasFailed >= 2:
		color.Yellow("Not bad!")
	default:
		color.Green("Today was a good day!")
	}
}
