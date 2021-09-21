package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
	log "github.com/sirupsen/logrus"
)

var lsusb_functional bool = false


func onContextError() {
	if !lsusb_functional {
		log.Warn("Unable to initialize USB discovery. Host might be incompatible with this peripheral manager. Trying again later...")
		time.Sleep(10 * time.Second)
		log.Info(string(debug.Stack()))
		os.Exit(0)
	}
}

func get_usb_context() *gousb.Context {
	defer onContextError()
	c := gousb.NewContext()
	lsusb_functional = true

	return c
}

func main() {
	log.Info("meetup: USB finder has started")

	ctx := get_usb_context()
	defer ctx.Close()

		devices, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
			identifier := fmt.Sprintf("%s:%s", desc.Vendor, desc.Product)
			vendor := usbid.Vendors[desc.Vendor]
			product := vendor.Product[desc.Product]

			description := fmt.Sprintf("USB device [%s] with ID %s. Protocol: %s",
				product,
				identifier,
				usbid.Classify(desc))

                        log.Info(description)

			return false
		})

	if err != nil {
		log.Error("Unable to discover USB devices")
	}

	log.Infof("Found %d USB devices", len(devices))
}
