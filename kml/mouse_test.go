package kml

import (
	"testing"
	"time"
)

func TestMouseMove(t *testing.T) {
	err := OpenDevice()
	if err != nil {
		t.Fatal(err)
	}
	MouseWheel(100)
	time.Sleep(400 * time.Millisecond)
	MouseWheel(-100)
}
