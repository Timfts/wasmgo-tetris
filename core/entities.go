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
