package core

const SPIN_CW = 1
const SPIN_CCW = -1

const BOARD_WIDTH = 10
const BOARD_HEIGHT = 20
const BOARD_OFFSET_X = 16
const BOARD_SIZE = BOARD_HEIGHT * BOARD_WIDTH

type PieceType string

const (
	L PieceType = "L"
	J PieceType = "J"
	Z PieceType = "Z"
	S PieceType = "S"
	O PieceType = "O"
	T PieceType = "T"
	I PieceType = "I"
)

var PIECE_COLORS = map[PieceType]int{
	"L": 0x0484d1,
	"J": 0xfb922b,
	"Z": 0xe53b44,
	"S": 0x63c64d,
	"O": 0xffe762,
	"T": 0xbf66bf,
	"I": 0x2ce8f4,
}

var PIECE_COORDS = map[PieceType][4]Point{
	"L": {
		{X: 1, Y: -1},
		{X: -1, Y: 0},
		{X: 0, Y: 0},
		{X: 1, Y: 0},
	},
	/*
		 	"J": 0xfb922b,
			"Z": 0xe53b44,
			"S": 0x63c64d,
			"O": 0xffe762,
			"T": 0xbf66bf,
			"I": 0x2ce8f4,
	*/
}

var PIECE_SPRITES = map[PieceType][8]byte{
	"L": {
		0b11111111,
		0b10000001,
		0b10111101,
		0b10100001,
		0b10100001,
		0b10100001,
		0b10000001,
		0b11111111,
	},
	"J": {
		0b11111111,
		0b10000001,
		0b10000101,
		0b10000101,
		0b10000101,
		0b10111101,
		0b10000001,
		0b11111111,
	},
	"Z": {
		0b11111111,
		0b10100101,
		0b11001001,
		0b10010011,
		0b10100101,
		0b11001001,
		0b10010011,
		0b11111111,
	},
	"S": {
		0b11111111,
		0b10100101,
		0b10010011,
		0b11001001,
		0b10100101,
		0b10010011,
		0b11001001,
		0b11111111,
	},
	"O": {
		0b11111111,
		0b10000001,
		0b10000001,
		0b10000001,
		0b10000001,
		0b10000001,
		0b10000001,
		0b11111111,
	},
	"T": {
		0b11111111,
		0b10000001,
		0b10111101,
		0b10100101,
		0b10100101,
		0b10111101,
		0b10000001,
		0b11111111,
	},
	"I": {
		0b11111111,
		0b10000001,
		0b10100101,
		0b10000001,
		0b10100101,
		0b10111101,
		0b10000001,
		0b11111111,
	},
}

var LEVEL_SPEED = []int{
	53,
	49,
	45,
	41,
	37,
	33,
	28,
	22,
	17,
	11,
	10,
	9,
	8,
	7,
	6,
	6,
	5,
	5,
	4,
	4,
	3,
}
