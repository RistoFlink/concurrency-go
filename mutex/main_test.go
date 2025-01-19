package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "hello, world"
	wg.Add(2)
	go updateMessage("XXX")
	go updateMessage("goodbye cruel world")
	wg.Wait()

	if msg != "goodbye cruel world" {
		t.Error("incorrect message")
	}
}
