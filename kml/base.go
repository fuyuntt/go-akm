package kml

import (
	"fmt"
	"github.com/karalabe/hid"
	"math/rand"
	"time"
)

const (
	vid              = 1155
	pid              = 22315
	controlInterface = 2
)

var device *hid.Device

func OpenDevice() error {
	deviceList := hid.Enumerate(vid, pid)
	var deviceInfo *hid.DeviceInfo
	for _, d := range deviceList {
		if d.Interface == controlInterface {
			deviceInfo = &d
			break
		}
	}
	if deviceInfo == nil {
		return fmt.Errorf("DeviceNotFound")
	}
	var err error
	device, err = deviceInfo.Open()
	if err != nil {
		return err
	}
	return nil
}

func WriteData(data []byte) error {
	if device == nil {
		return DeviceNotOpen
	}
	_, err := device.Write(data)
	return err
}

func randSleep() {
	r := 50 + rand.Intn(20)
	time.Sleep(time.Duration(r) * time.Millisecond)
}
