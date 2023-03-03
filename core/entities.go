package core

type GameState struct {
	gameOver bool
	level    int
	score    int
	best     int
	lines    int
}
type Point struct {
	X int
	Y int
}

type Piece struct {
	x      uint8
	y      uint8
	kind   PieceType
	coords [8]int8
}
