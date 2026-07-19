package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bdstark/ble"
	"github.com/bdstark/ble-examples/lib/dev"
)

func main() {
	macAddr := flag.String("addr", "", "peripheral MAC address")
	flag.Parse()
	hciDevice, err := dev.NewDevice("default")
	if err != nil {
		panic(err)
	}
	ble.SetDefaultDevice(hciDevice)

	filter := func(a ble.Advertisement) bool {
		return strings.ToUpper(a.Addr().String()) == strings.ToUpper(*macAddr)
	}

	// Scan for device
	log.Printf("Scanning for %s\n", *macAddr)
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), time.Second*300))
	client, err := ble.Connect(ctx, filter)
	if err != nil {
		panic(err)
	}

	for {
		rssi, err := client.ReadRSSI(context.Background())
		if err != nil {
			fmt.Printf("Client side RSSI: error: %v\n", err)
		} else {
			fmt.Printf("Client side RSSI: %d\n", rssi)
		}
		time.Sleep(time.Second)
	}

}
