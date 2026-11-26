//go:build windows

package controller

import (
	"errors"
	"syscall"
	"unsafe"
)

const (
	up    = 0x0001
	down  = 0x0002
	left  = 0x0004
	right = 0x0008

	lb = 0x0100
	rb = 0x0200

	l3 = 0x0040
	r3 = 0x0080

	a = 0x1000
	b = 0x2000
	x = 0x4000
	y = 0x8000

	triggerThreshold = 30
)

type gamepad struct {
	Buttons      uint16
	LeftTrigger  byte
	RightTrigger byte
	LeftX        int16
	LeftY        int16
	RightX       int16
	RightY       int16
}

type xinputState struct {
	PacketNumber uint32
	Gamepad      gamepad
}

type Reader struct {
	getState *syscall.LazyProc
	index    uintptr
}

func NewReader(index uint) (*Reader, error) {
	for _, name := range []string{
		"xinput1_4.dll",
		"xinput9_1_0.dll",
		"xinput1_3.dll",
	} {
		proc := syscall.NewLazyDLL(name).NewProc("XInputGetState")

		if proc.Find() == nil {
			return &Reader{
				getState: proc,
				index:    uintptr(index),
			}, nil
		}
	}

	return nil, errors.New("XInputGetState unavailable")
}

func (r *Reader) Read() State {
	s := Empty()

	var raw xinputState

	status, _, _ := r.getState.Call(
		r.index,
		uintptr(unsafe.Pointer(&raw)),
	)

	if status != 0 {
		return s
	}

	s.Connected = true

	bits := raw.Gamepad.Buttons
	m := s.Buttons

	// Directional inputs
	m["up"] = bits&up != 0
	m["down"] = bits&down != 0
	m["left"] = bits&left != 0
	m["right"] = bits&right != 0

	// Punch buttons
	m["1p"] = bits&x != 0
	m["2p"] = bits&y != 0
	m["3p"] = bits&rb != 0
	m["4p"] = bits&lb != 0

	// Kick buttons
	m["1k"] = bits&a != 0
	m["2k"] = bits&b != 0

	// Analog triggers represented as digital inputs
	m["3k"] = raw.Gamepad.RightTrigger > triggerThreshold
	m["4k"] = raw.Gamepad.LeftTrigger > triggerThreshold

	// Stick clicks (including macro aliases)
	m["l3"] = bits&l3 != 0
	m["r3"] = bits&r3 != 0

	return s
}
