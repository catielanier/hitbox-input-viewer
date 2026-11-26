//go:build !windows

package controller

import "errors"

type Reader struct{}

func NewReader(index uint) (*Reader, error) {
	return nil, errors.New(
		"XInput collector only runs on Windows; " +
			"cross-compile with GOOS=windows",
	)
}

func (*Reader) Read() State {
	return Empty()
}
