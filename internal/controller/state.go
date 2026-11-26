package controller

// State is the normalized controller state. It deliberately represents logical
// XInput controls, not physical GPIO switches inside the controller firmware.
type State struct {
	Connected bool            `json:"connected"`
	Buttons   map[string]bool `json:"buttons"`
}

var labels = []string{
	"up", "down", "left", "right",
	"1p", "2p", "3p", "4p",
	"1k", "2k", "3k", "4k",
	"l3", "r3",
}

func Empty() State {
	buttons := make(map[string]bool, len(labels))

	for _, label := range labels {
		buttons[label] = false
	}

	return State{Buttons: buttons}
}

func Equal(a, b State) bool {
	if a.Connected != b.Connected {
		return false
	}

	for _, k := range labels {
		if a.Buttons[k] != b.Buttons[k] {
			return false
		}
	}

	return true
}
