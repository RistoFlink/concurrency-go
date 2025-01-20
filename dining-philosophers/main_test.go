package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("orderFinished length is %d, expected 5", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"Zero delay", 0 * time.Second},
		{"Quarter second delay", 250 * time.Millisecond},
		{"Half second delay", 500 * time.Millisecond},
	}

	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		thinkTime = e.delay
		sleepTime = e.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: orderFinished length is %d, expected 5", e.name, len(orderFinished))
		}
	}
}
