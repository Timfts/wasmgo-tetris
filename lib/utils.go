package lib

import "cart/w4"

func SetupColors() {
	// generate palete by https://lospec.com/palette-list
	w4.PALETTE[0] = 0xd8dee9
	w4.PALETTE[2] = 0x4c566a
	w4.PALETTE[3] = 0x2e3440
}

var smiley = [8]byte{
	0b11000011,
	0b10000001,
	0b00100100,
	0b00100100,
	0b00000000,
	0b00100100,
	0b10011001,
	0b11000011,
}

func DrawSmile() {
	w4.Blit(&smiley[0], 76, 76, 8, 8, w4.BLIT_1BPP)
}
