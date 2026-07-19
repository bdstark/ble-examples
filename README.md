# ble-examples

Example programs for [bdstark/ble](https://github.com/bdstark/ble), a
maintained fork of [go-ble/ble](https://github.com/go-ble/ble) (module path
`github.com/go-ble/ble`).

These lived in the library repo as `examples/` until 2026-07; they were moved
here so the library carries no example-only dependencies (`urfave/cli`).

- `basic/scanner` — scan and print advertisements
- `basic/advertiser` — advertise name and services
- `basic/explorer` — connect and dump a peripheral's GATT profile
- `basic/rssi` — read RSSI from a peer
- `basic/server` — a GATT server with count/echo/battery characteristics
- `blesh` — interactive BLE shell
- `lib`, `lib/dev` — shared helpers (default device per platform)

The `go.mod` `replace` directive points `github.com/go-ble/ble` at the fork's
`main`; run `go get github.com/go-ble/ble@main && go mod tidy` to track it.
