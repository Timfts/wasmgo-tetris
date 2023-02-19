package main

import "cart/src/w4"

func setupColors() {
	// generate palete by https://lospec.com/palette-list
	w4.PALETTE[0] = 0x8cad28
	w4.PALETTE[1] = 0x6c9421
	w4.PALETTE[2] = 0x426b29
	w4.PALETTE[3] = 0x214231
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

func drawSmile() {
	w4.Blit(&smiley[0], 76, 76, 8, 8, w4.BLIT_1BPP)
}
