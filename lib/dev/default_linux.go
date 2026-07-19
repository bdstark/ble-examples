package dev

import (
	"github.com/bdstark/ble"
	"github.com/bdstark/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
