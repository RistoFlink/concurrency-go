package main

import (
	"time"

	"github.com/fatih/color"
)

// variables
var (
	seatingCapacity = 10
	arrivalRate     = 100
	cutDuration     = 1000 * time.Millisecond
	timeOpen        = 10 * time.Second
)

func main() {
	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels if needed
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barber shop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HaircutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for business!")

	// add barbers

	// start the barber shop (as a goroutine)

	// add clients

	// block (= keep application going) until the barber shop is closed
}
