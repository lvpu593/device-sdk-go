package main

import "github.com/lvpu593/device-sdk-go/example/driver"

const (
	serviceName = "device-simple"
)

func main() {

	sd := driver.SimpleDriver{}
	startup.Bootstrap(serviceName, device.Version, &sd)
}
