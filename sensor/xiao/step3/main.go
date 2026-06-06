package main

import (
	"machine"
	"time"
)

var (
	green = machine.D1
	red   = machine.D8

	button = machine.D10
)

func main() {
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		if button.Get() {
			green.High()
			red.Low()
		} else {
			green.Low()
			red.High()
		}

		time.Sleep(time.Millisecond * 100)
	}
}
