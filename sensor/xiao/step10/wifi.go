package main

import (
	"net/netip"
	"time"

	"tinygo.org/x/drivers/netdev"
	"tinygo.org/x/espradio"
	nl "tinygo.org/x/espradio/netlink"
)

var (
	ssid     string
	password string
	port     string = ":80"

	link nl.Esplink
)

const apIP = "192.168.4.1"

func connectWifi() {
	// wait a bit for serial
	time.Sleep(2 * time.Second)

	link.ArenaPoolSize = 1024 * 42
	netdev.UseNetdev(&link)

	println("Connecting to WiFi...")
	err := link.NetConnectAP(nl.APConnectParams{
		APConfig: espradio.APConfig{
			SSID:     ssid,
			Password: password,
			Channel:  6,
		},
		StaticAddr:       netip.MustParseAddr(apIP),
		EnableDHCPServer: true,
		MaxUDPPorts:      2,
		MaxTCPPorts:      4,
		PassivePeers:     8,
	})

	if err != nil {
		failMessage("could not connect to WiFi: " + err.Error())
	}
}

func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}
