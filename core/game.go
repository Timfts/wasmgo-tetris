package core

import (
	"cart/w4"
	"math/rand"
)

type Game struct {
	state               GameState
	previousGamepad     byte
	board               [BOARD_SIZE]string
	clearAnimationDelay byte
	piece               Piece
}

func NewGame() Game {
	initialState := GameState{
		gameOver: false,
	}

	firstPieceType := PieceType(rand.Intn(7))
	pieceCoords := getPieceCoords(firstPieceType)

	firstPiece := Piece{
		x:      5,
		y:      1,
		kind:   firstPieceType,
		coords: pieceCoords,
	}

	return Game{
		state: initialState,
		piece: firstPiece,
	}
}

func (s *Game) Setup() {
	w4.PALETTE[0] = 0xd8dee9
	w4.PALETTE[1] = 0x4c566a
	w4.PALETTE[2] = 0x4c566a
	w4.PALETTE[3] = 0x2e3440
}

func (s *Game) Update() {
	coords := getPieceCoords(J)

	logValue(coords[7])

	/* 	ui8ToString := func(value uint8) string {
	   		return strconv.FormatUint(uint64(value), 10)
	   	}

	   	var gamepad = *w4.GAMEPAD1

	   	w4.Trace(ui8ToString((gamepad)))
	   	var pressedPreviousAndCurrent = gamepad ^ s.previousGamepad
	   	var onlyPressedThisFrame = gamepad & pressedPreviousAndCurrent

	   	s.previousGamepad = gamepad

	   	if onlyPressedThisFrame&w4.BUTTON_LEFT != 0 {
	   		w4.Trace("pressed left")
	   	} */

	s.drawPiece(s.piece.x, s.piece.y, s.piece.kind, s.piece.coords)
	s.drawGUIBg()
	s.drawStats()

	if s.clearAnimationDelay != 0 && !s.state.gameOver {
		*w4.DRAW_COLORS = 0x42
		// draw piece
	}

}

func (s *Game) restart() {
	s.state.gameOver = false
}

func (s *Game) drawStats() {
	*w4.DRAW_COLORS = 4
	w4.Text("LEVEL", 104, 16)
	w4.Text("SCORE", 104, 40)
	w4.Text("BEST", 104, 64)

	*w4.DRAW_COLORS = 1
	w4.Text("2", 104, 24+2)
	w4.Text("3", 104, 48+2)
	w4.Text("4", 104, 72+2)
}

func (s *Game) drawGUIBg() {
	*w4.DRAW_COLORS = 3
	w4.Rect(0, 0, BOARD_OFFSET_X, 160)
	w4.Rect(BOARD_OFFSET_X+BOARD_WIDTH*8, 0, 160-(BOARD_OFFSET_X+BOARD_WIDTH*8), 160)

	*w4.DRAW_COLORS = 4
	w4.Rect(BOARD_OFFSET_X, 0, 1, 160)
	w4.Rect(BOARD_OFFSET_X+BOARD_WIDTH*8-1, 0, 1, 160)
}

func (s *Game) drawGameoverText() {
	*w4.DRAW_COLORS = 4
	w4.Text("GAME OVER", BOARD_OFFSET_X+4, 64)
	w4.Text(" PRESS X", BOARD_OFFSET_X+4, 64+8+2)
}

func (s *Game) startNewRound() {

}

func (s *Game) drawBlock(pieceType PieceType, x int, y int, disabled bool) {
	*w4.DRAW_COLORS = 4
	spriteToUse := PIECE_SPRITES[pieceType]
	w4.Blit(&spriteToUse[0], BOARD_OFFSET_X+x, y, 8, 8, w4.BLIT_1BPP)
}

func (s *Game) drawPiece(x uint8, y uint8, pieceType PieceType, pieceCoords [8]int8) {
	c := 0
	for n := 0; n < 4; n++ {
		cx := int(pieceCoords[c])
		cy := int(pieceCoords[c+1])
		s.drawBlock(pieceType, 8*(int(x)+cx), 8*(int(y)+cy), false)
		c += 2
	}
}
