package kml

import (
	"testing"
)

func Test_Device(t *testing.T) {
	err := OpenDevice()
	if err != nil {
		t.Fatal(err)
		return
	}
}
