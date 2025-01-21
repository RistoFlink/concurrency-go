package main

import (
	"fmt"
	"math/rand"
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
	shop.addBarber("Dennis")
	shop.addBarber("Frank")
	shop.addBarber("Dee")
	shop.addBarber("Mac")
	shop.addBarber("Charlier")
	shop.addBarber("Cricket")

	// start the barber shop (as a goroutine)
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen) // blocks until timeOpen passes
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()
	// add clients
	i := 1

	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()
	// block (= keep application going) until the barber shop is closed
	<-closed
}
