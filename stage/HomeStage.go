package stage

import (
	"LonelyChild/game"
	"LonelyChild/gobject"
	"os"
)

func GetHomeStage() *game.Stage {
	homeStage := new(game.Stage)
	homeStage.Start = startHomeStage
	return homeStage
}

func startHomeStage(s *game.Stage) {
	LCgame := s.LCgame
	player := s.LCgame.Data.Player
	for (player.Position < 300 && player.Position >= 200) || (player.Position >= 700 && player.Position < 800) {
		switch player.Position {
		case 200:
			zeroho()
			break
		case 201:
			oneho()
			break
		case 202:
			twoho()
			break

		case 700:
			gzeroho()
			break
		case 701:
			goneho()
			break
		case 702:
			gtwoho()
			break
		case 703:
			gthreeho()
			break
		}
		LCgame.SaveGameData()
	}
}

func zeroho() {
	game.WriteS("I am back home...\n" +
		"I can't go there.\n" +
		"They hate me.\n" +
		"It will all start from the begining.\n" +
		"Mother: " + player.Name + "!!!\n" +
		"Mother: You are back...\n")
	game.GetEnter()
	player.Position = (1)
}

func oneho() {
	if LCgame.SetStage(GetFightStage(gobject.GetMother(false))) == 1 {
		player.Kill(16)
		game.WriteS("You killed you mother... Do you feel better?")
		player.Position = (202)
	} else {
		game.WriteS("Mother: ...\n" +
			"Mother: I am sorry how i used to be.\n" +
			"Mother: I will change.\n" +
			"Mother: Let us go to Father.")
		player.Position = (202)
	}
}

func twoho() {
	if player.GetKilled(16) {
		game.WriteS("You go into your house.\n" +
			"You see your dad.\n" +
			"Your dad sees blood you hands.\n" +
			"Your dad attacks you.\n")
		LCgame.SetStage(GetFightStage(gobject.GetFather(true)))
		game.WriteS("Your siblings saw that and flew.\n" +
			"You feel great and relax a little bit.\n")
		player.Position = (12345)

	} else {
		game.WriteS("You follow your mom.\n" +
			"She holds your hand.\n" +
			"She talks now with you Father.\n" +
			"You Father attacks you.\n")
		if LCgame.SetStage(GetFightStage(gobject.GetFather(false))) == 1 {
			player.GetKilled(15)
			game.WriteS("What did you do?...")
			game.WriteS("This was your Father!")
			game.WriteS("Your Mother attacks you.")
			LCgame.SetStage(GetFightStage(gobject.GetMother(true)))
			game.WriteS("Your siblings saw that and flew.\n" +
				"You feel great and relax a little bit.\n")
			player.Position = (12345)
		} else {
			game.WriteS("Parents: We are so sorry to you...\n" +
				"Father: I can't understand why i hated you...\n" +
				"Parents: We will talk with your siblings about that.\n" +
				"They talk with you siblings\n" +
				"You and your family are now together")
			player.Position = (1234)
		}

	}
	os.Exit(0)
}

func gzeroho() {
	game.WriteS("You reached your home town.\n" +
		"You see some Pupils that bullied you.\n")
	if player.ConsistsFlowey {
		game.WriteS("Let's kill them.")
	}
	if LCgame.SetStage(GetFightStage(gobject.GetPupil())) == 1 {
		player.Kill(10)
		game.WriteS("The ground under the dead body disapeared, he fell down\n")
	} else {
		game.WriteS("The ground under the pupil disapeared, he fell down.\n" +
			"Now he is dead.\n")
	}
	game.WriteS("You hear a voice whispers: I am coming back\n")

	game.WriteS(player.Name + ": one less")
	player.Position = (701)
}

func goneho() {
	game.WriteS("The teacher looks at you.\n" +
		"He wants to scream.\n" +
		"You attacked him.\n")
	if LCgame.SetStage(GetFightStage(gobject.GetTeacher())) == 1 {
		player.Kill(11)
		game.WriteS("A raven comes at pick his eye off.\n")
	} else {
		game.WriteS("A swarm of raven comes and try to eat your teacher.\n" +
			"Now he is dead.\n")
	}

	game.WriteS("You hear a voice whispers: I am not far away.\n")

	game.WriteS(player.Name + ": one less")
	player.Position = (702)

}

func gtwoho() {
	game.WriteS("The Principal looks at you.\n" +
		"He comes right at you.\n" +
		"Principal: You little bas.\n" +
		"You attack him.\n")
	if LCgame.SetStage(GetFightStage(gobject.GetPrincipal())) == 1 {
		player.Kill(12)
		game.WriteS("His body starts to burn.\n")
	} else {
		game.WriteS("His body catched fire.\nHe screams.\n" +
			"Now he is dead.\n")
	}

	game.WriteS("You hear a voice whispers: Go home and let it end\n")

	game.WriteS(player.Name + ": one less")
	player.Position = (703)
}

func gthreeho() {
	game.WriteS("You go the way to your \"home\".\n" +
		"When you reached you see you whole family\n" +
		"They want to say something.\n")
	LCgame.SetStage(GetFightStage(gobject.GetBrother(true)))
	LCgame.SetStage(GetFightStage(gobject.GetSister(true)))
	LCgame.SetStage(GetFightStage(gobject.GetFather(true)))
	LCgame.SetStage(GetFightStage(gobject.GetMother(true)))
	player.Kill(13)
	player.Kill(14)
	player.Kill(15)
	player.Kill(16)

	game.WriteS("You killed them before they said a word:\n\n" +
		"Their body starts to melt.\n" +
		"You feel happy.\n")

	game.WriteS("You hear a voice whispers: Now i will take you.\n")

	player.Position = (800)
}
