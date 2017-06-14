package main

import (
	"LonelyChild/game"
	"LonelyChild/stage"
)

func main() {
	LCgame := game.NewGame()
	LCgame.SetStage(stage.GetBootStage())
}
