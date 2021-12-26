package kml

import "sync"

var kbLock sync.Mutex
var keyState = []uint8{1, 0, 0, 0, 0, 0, 0, 0, 0}

func PressKey(vk Vk) error {
	kbLock.Lock()
	defer kbLock.Unlock()
	if vk >= VkLCtrl {
		keyState[1] |= 1 << (vk - VkLCtrl)
	} else {
		for i := 3; i < 9; i++ {
			if keyState[i] == 0 {
				keyState[i] = uint8(vk)
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
		var move bool
		for i := 3; i < 9; i++ {
			if keyState[i] == 0 {
				break
			} else if move {
				keyState[i-1] = keyState[i]
			} else if keyState[i] == uint8(vk) {
				keyState[i] = 0
				move = true
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
