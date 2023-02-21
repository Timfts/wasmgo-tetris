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
	state GameState
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
	s.Setup()
	s.DrawStats()
}

func (s *Game) restart() {

	s.state.gameOver = false
}

func (s *Game) DrawStats() {
	/* *w4.DRAW_COLORS = 4 */
	w4.Text("LEVEL", 104, 16)
	w4.Text("SCORE", 104, 40)
	w4.Text("BEST", 104, 64)

	/* *w4.DRAW_COLORS = 1 */
	w4.Text("2", 104, 24+2)
	w4.Text("3", 104, 48+2)
	w4.Text("4", 104, 72+2)
}
