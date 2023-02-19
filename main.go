package main

import (
	"cart/lib"
	"cart/w4"
)

// #region gameState
var gameover = false

// #endregion

func start() {
	lib.SetupColors()
}

//go:export update
func update() {
	lib.SetupColors()
	*w4.DRAW_COLORS = 2
	w4.Text("potato", 10, 10)

	var gamepad = *w4.GAMEPAD1
	if gamepad&w4.BUTTON_1 != 0 {
		*w4.DRAW_COLORS = 4
	}

	lib.DrawSmile()
	w4.Text("Press X to blink", 16, 90)

}
