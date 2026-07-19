package dev

import (
	"github.com/bdstark/ble"
	"github.com/bdstark/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
