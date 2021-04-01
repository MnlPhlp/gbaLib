package gbaLib

type Button struct {
	KeyCode uint16
}

var Buttons = struct {
	Up, Down, Right, Left Button
	L, R                  Button
	A, B                  Button
	Start, Select         Button
}{
	A:      Button{KeyCode: 1 << 0},
	B:      Button{KeyCode: 1 << 1},
	Select: Button{KeyCode: 1 << 2},
	Start:  Button{KeyCode: 1 << 3},
	Right:  Button{KeyCode: 1 << 4},
	Left:   Button{KeyCode: 1 << 5},
	Up:     Button{KeyCode: 1 << 6},
	Down:   Button{KeyCode: 1 << 7},
	R:      Button{KeyCode: 1 << 8},
	L:      Button{KeyCode: 1 << 9},
}

func (button *Button) IsPressed() bool {
	// inverted because the key bits are active low
	return !Register.Key.KeyPad.HasBits(uint16(button.KeyCode))
}
