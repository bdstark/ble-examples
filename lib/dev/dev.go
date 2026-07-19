package dev

import (
	"github.com/bdstark/ble"
)

// NewDevice ...
func NewDevice(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}
