package stage

import "LonelyChild/game"
import "LonelyChild/gobject"

func GetFightStage(enemy gobject.GObject) *game.Stage {
	fightStage := new(game.Stage)
	fightStage.Start = startFightStage
	return fightStage
}

func startFightStage(s *game.Stage) {
	s.Happened = 0
}
