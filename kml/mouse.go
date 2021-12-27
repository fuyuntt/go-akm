package kml

import "sync"

var mouseLock sync.Mutex
var mouseState = []byte{2, 0, 0, 0, 0, 0, 0, 0, 0}

func PressButton(button MButton) error {
	mouseLock.Lock()
	defer mouseLock.Unlock()
	mouseState[1] |= 1 << button
	return WriteData(mouseState)
}
func ReleaseButton(button MButton) error {
	mouseLock.Lock()
	defer mouseLock.Unlock()
	mouseState[1] &= ^(1 << button)
	return WriteData(mouseState)
}

func ClickButton(button MButton) error {
	err := PressButton(button)
	if err != nil {
		return err
	}
	randSleep()
	return ReleaseButton(button)
}

func MouseMove(x, y int8) error {
	mouseLock.Lock()
	defer mouseLock.Unlock()
	mouseState[2] = byte(x)
	mouseState[3] = byte(y)
	err := WriteData(mouseState)
	mouseState[2] = 0
	mouseState[3] = 0
	return err
}

func MouseWheel(w int8) error {
	mouseLock.Lock()
	defer mouseLock.Unlock()
	mouseState[4] = byte(w)
	err := WriteData(mouseState)
	mouseState[4] = 0
	return err
}
