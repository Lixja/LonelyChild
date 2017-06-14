package main

import (
	"LonelyChild/game"
	"LonelyChild/stage"
)

func main() {
	var lcgame game.Game
	fstage := stage.GetBootStage()
	lcgame.SetStage(fstage)
}
