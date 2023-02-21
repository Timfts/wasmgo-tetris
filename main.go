package main

import (
	"cart/core"
)

var gameInstance = core.NewGame()

//go:export start
func start() {
	gameInstance.Setup()
}

//go:export update
func update() {

	gameInstance.Update()

}
