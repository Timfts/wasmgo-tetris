package core

import "cart/w4"

type GameState struct {
	gameOver bool
	level    int
	score    int
	best     int
	lines    int
}

type Game struct {
	state           GameState
	previousGamepad byte
}

func NewGame() Game {
	// setup initial state
	initialState := GameState{
		gameOver: false,
	}

	return Game{
		state: initialState,
	}
}

func (s *Game) Setup() {
	w4.PALETTE[0] = 0xd8dee9
	w4.PALETTE[2] = 0x4c566a
	w4.PALETTE[3] = 0x2e3440
}

func (s *Game) Update() {
	var gamepad = *w4.GAMEPAD1
	var pressedThisFrame = gamepad & (gamepad ^ s.previousGamepad)
	s.previousGamepad = gamepad

	if pressedThisFrame&w4.BUTTON_LEFT != 0 {

	}

	s.drawGUIBg()
	s.drawStats()
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
