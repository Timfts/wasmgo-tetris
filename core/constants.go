package core

const SPIN_CW = 1
const SPIN_CCW = -1

const BOARD_WIDTH = 10
const BOARD_HEIGHT = 20
const BOARD_OFFSET_X = 16
const BOARD_SIZE = BOARD_HEIGHT * BOARD_WIDTH

type PieceType int8

const (
	L PieceType = 0
	J PieceType = 1
	Z PieceType = 2
	S PieceType = 3
	O PieceType = 4
	T PieceType = 5
	I PieceType = 6
)

var PIECE_COLORS = []int{
	0x0484d1, // L: blue
	0xfb922b, // J: orange
	0xe53b44, // Z: red
	0x63c64d, // S: green
	0xffe762, // O: yellow
	0xbf66bf, // T: purple
	0x2ce8f4, // I: cyan
}

var PIECE_COORDS = []int8{
	// L
	1, -1,
	-1, 0,
	0, 0,
	1, 0,

	// J
	-1, -1,
	-1, 0,
	0, 0,
	1, 0,

	// Z
	0, -1,
	-1, -1,
	0, 0,
	1, 0,

	// S
	0, -1,
	1, -1,
	0, 0,
	-1, 0,

	// O
	-1, -1,
	-1, 0,
	0, 0,
	0, -1,

	// T
	0, -1,
	-1, 0,
	0, 0,
	1, 0,

	// I
	-2, 0,
	-1, 0,
	0, 0,
	1, 0,
}

var PIECE_SPRITES = [][8]uint8{
	{
		0b11111111,
		0b10000001,
		0b10111101,
		0b10100001,
		0b10100001,
		0b10100001,
		0b10000001,
		0b11111111,
	},
	{
		0b11111111,
		0b10000001,
		0b10000101,
		0b10000101,
		0b10000101,
		0b10111101,
		0b10000001,
		0b11111111,
	},
	{
		0b11111111,
		0b10100101,
		0b11001001,
		0b10010011,
		0b10100101,
		0b11001001,
		0b10010011,
		0b11111111,
	},
	{
		0b11111111,
		0b10100101,
		0b10010011,
		0b11001001,
		0b10100101,
		0b10010011,
		0b11001001,
		0b11111111,
	},
	{
		0b11111111,
		0b10000001,
		0b10000001,
		0b10000001,
		0b10000001,
		0b10000001,
		0b10000001,
		0b11111111,
	},
	{
		0b11111111,
		0b10000001,
		0b10111101,
		0b10100101,
		0b10100101,
		0b10111101,
		0b10000001,
		0b11111111,
	},
	{
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

var LEVEL_SPEED = []int8{
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
