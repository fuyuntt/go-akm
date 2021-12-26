package kml

import (
	"testing"
)

func Test_Kb(t *testing.T) {
	err := OpenDevice()
	if err != nil {
		t.Fatal(err)
	}
	err = PressKey(VkLCtrl)
	if err != nil {
		t.Fatal(err)
	}
	err = PressAndReleaseKey(VkC)
	if err != nil {
		t.Fatal(err)
	}
	err = PressAndReleaseKey(VkV)
	if err != nil {
		t.Fatal(err)
	}
	err = PressAndReleaseKey(VkV)
	if err != nil {
		t.Fatal(err)
	}
	err = ReleaseKey(VkLCtrl)
	if err != nil {
		t.Fatal(err)
	}

}
