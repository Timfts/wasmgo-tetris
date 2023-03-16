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
	case string:
		return value.(string)
	default:
		return ""
	}
}

func logValue(value interface{}) {
	w4.Trace(convertToString(value))
}
