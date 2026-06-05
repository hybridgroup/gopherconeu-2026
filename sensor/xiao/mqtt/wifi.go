package main

import (
	"time"

	"tinygo.org/x/drivers/netdev"
	nl "tinygo.org/x/drivers/netlink"
	link "tinygo.org/x/espradio/netlink"
)

var (
	ssid     string
	password string
)

func connectWifi() {
	// wait a bit for serial
	time.Sleep(2 * time.Second)

	link := link.Esplink{}
	netdev.UseNetdev(&link)

	println("Connecting to WiFi...")
	err := link.NetConnect(&nl.ConnectParams{
		Ssid:       ssid,
		Passphrase: password,
	})

	if err != nil {
		failMessage("could not connect to WiFi: " + err.Error())
	}
}
