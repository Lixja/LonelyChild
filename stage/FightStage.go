package stage

import "LonelyChild/game"

func GetFightStage(GObject enemy) *game.Stage {
	fightStage := new(game.Stage)
	fightStage.Start = startFightStage
	return fightStage
}

func startFightStage(s *game.Stage) {
	happened = 0
}
