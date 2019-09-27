package main

import (
	"github.com/lvpu593/device-sdk-go/example/driver"
	"github.com/lvpu593/device-sdk-go/pkg/startup"
)

const (
	serviceName = "device-simple"
	version = "0.0.1"
)

func main() {

	sd := driver.SimpleDriver{}
	startup.BootStrap(serviceName, version, &sd)
}
