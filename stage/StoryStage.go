package stage

import "LonelyChild/game"
import "os"

func GetStoryStage() *game.Stage {
	storyStage := new(game.Stage)
	storyStage.Start = startStoryStage
	return storyStage
}

func startStoryStage(s *game.Stage) {
	player := &s.LCgame.Data.Player
	for true {
		if (player.Position < 100 && player.Position >= 0) || (player.Position >= 500 && player.Position < 600) {
			s.LCgame.SetStage(GetForestStage())
			askForPause(s)
		}
		if (player.Position < 200 && player.Position >= 100) || (player.Position >= 600 && player.Position < 700) {
			//game.SetStage(new HouseOfHorrorsStage());
			askForPause(s)
		}
		if (player.Position < 300 && player.Position >= 200) || (player.Position >= 700 && player.Position < 800) {
			//game.SetStage(new HomeStage());
			askForPause(s)
		}
		if player.Position == 800 || player.Position == 999 {
			//game.SetStage(new WorldStage());
		}
	}
}

func askForPause(s *game.Stage) {
	s.LCgame.SaveGameData()
	game.WriteWall()
	game.WriteWall()
	answer := game.GetInputWithQuestionYesNo("Do you wanna make a break?")
	if answer {
		os.Exit(0)
	}
	game.WriteWall()
	game.WriteWall()
}
