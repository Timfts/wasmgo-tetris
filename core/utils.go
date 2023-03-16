package core

import (
	"cart/w4"
	"strconv"
)

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

func getPieceCoords(pieceType PieceType) [8]int8 {
	var coords [8]int8
	for n := 0; n < 8; n++ {
		firstItemIndex := pieceType * 8
		itemPointer := &PIECE_COORDS[int8(firstItemIndex)+int8(n)]
		coords[n] = *itemPointer
	}
	return coords
}

func convertToString(value interface{}) string {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case int8:
		return strconv.Itoa(int(value.(int8)))
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case bool:
		return strconv.FormatBool(value.(bool))
	default:
		return ""
	}
}

func logValue(value interface{}) {
	w4.Trace(convertToString(value))
}
