package kml

import "sync"

var kbLock sync.Mutex
var keyState = []byte{1, 0, 0, 0, 0, 0, 0, 0, 0}

func PressKey(vk Vk) error {
	kbLock.Lock()
	defer kbLock.Unlock()
	if vk >= VkLCtrl {
		keyState[1] |= 1 << (vk - VkLCtrl)
	} else {
		for i := 3; i < 9; i++ {
			if keyState[i] == 0 || keyState[i] == byte(vk) {
				keyState[i] = byte(vk)
				break
			}
		}
	}
	return WriteData(keyState)
}

func ReleaseKey(vk Vk) error {
	kbLock.Lock()
	defer kbLock.Unlock()
	if vk >= VkLCtrl {
		keyState[1] &= ^(1 << (vk - VkLCtrl))
	} else {
		for i := 3; i < 9; i++ {
			if keyState[i] == byte(vk) {
				keyState[i] = 0
			}
		}
	}
	return WriteData(keyState)
}

func PressAndReleaseKey(vk Vk) error {
	err := PressKey(vk)
	if err != nil {
		return err
	}
	randSleep()
	return ReleaseKey(vk)
}
