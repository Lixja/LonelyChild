package stage

import (
	"LonelyChild/game"
	"LonelyChild/gobject"
	"math/rand"
	"strconv"
	"strings"
)

func GetForestStage() *game.Stage {
	forestStage := new(game.Stage)
	forestStage.Start = startForestStage
	return forestStage
}

func startForestStage(s *game.Stage) {
	LCgame = s.LCgame
	player = &LCgame.Data.Player
	for (player.Position < 100 && player.Position >= 0) || (player.Position >= 500 && player.Position < 600) {
		switch player.Position {
		case 0:
			zerof()
			break
		case 1:
			onef()
			break
		case 2:
			twof()
			break
		case 3:
			threef()
			break
		case 4:
			fourf()
			break
		case 5:
			fivef()
			break
		case 6:
			sixf()
			break
		case 7:
			sevenf()
			break
		case 8:
			eightf()
			break
		case 502:
			gtwof()
			break
		case 503:
			gthreef()
			break
		case 504:
			gfourf()
			break
		case 505:
			gfivef()
			break
		case 506:
			gsixf()
			break
		case 507:
			gsevenf()
			break
		case 508:
			geightf()
			break
		case 509:
			gninef()
		}
		LCgame.SaveGameData()
	}
}

func zerof() {
	game.WriteS("You look around.")
	game.Sleep(1500)
	game.WriteS("You don't know were you are.\n" +
		"You can't even remember the way or since when you were in a forest.\n" +
		"But you think it's safe now\n")
	game.WritelnSlow(". . .", 500)
	game.WriteS("You think about your family and the human race.\n" +
		"You hate them.\n" +
		"Everybody\n\n" +
		"You start laughing with no reason.\n" +
		"You think you get crazy.\n" +
		"After all ")
	game.Sleep(500)
	game.WritelnSlow("you hear a stone talking?", 120)
	player.Position = 1
}

func onef() {
	if LCgame.SetStage(GetFightStage(gobject.GetMasterStone())) == 1 {
		player.Kill(0)
		player.Position = 502
	} else {
		player.Position = 2
	}
}

func twof() {
	game.WriteS("No.\n")
	game.Sleep(500)
	game.WriteS("You feel great for not beeing an asshole.\n")
	game.GetEnter()
	game.WriteS("You ask the stone why you can hear him.\n\n" +
		"The stone tells you you are in the forest of dreams.\n" +
		"The only place on earth that can be enter by a poor soul.\n\n" +
		"You ask if somebody else is also on this place.\n\n" +
		"The stone tells you to follow him.\n\n" +
		"So you do\n\n" +
		"The stone introduces you Snake.\n\n\n" +
		"Snake: Why do you bring a human beeing to our place\n" +
		"Stone: He seemed friendly. He did not attack me.\n" +
		"Snake: He tricked you! A human can't be friendly!\n")
	player.Position = 3
}

func threef() {
	choice := false
	if LCgame.SetStage(GetFightStage(gobject.GetSnake())) == 1 {
		player.Kill(1)
		game.WriteS("Stone: Why did you do that?")
		choice = game.GetInputWithQuestionYesNo("Do you wanna fight Stone?")
		if choice {
			if LCgame.SetStage(GetFightStage(gobject.GetMasterStone())) == 1 {
				player.Kill(0)
				player.Position = 504
				return
			} else {
				game.WriteS("Stone flew.\n")
			}
			game.WriteS("Stone flew.\n")
		}
		player.Position = 504
	} else {
		player.Position = 4
	}
}

func fourf() {
	game.WriteS("Snake: ...\n" +
		"Snake: I am sorry I attacked you.\n" +
		"Snake: Every human i saw until now was bad.\n" +
		"Snake: But you seem so different.\n")
	game.Sleep(500)
	game.WriteS("Snake: You are welcome to stay at this place.\n" +
		"Snake: But be careful, most beeings would attack you if they see you.\n" +
		"Snake: We hate humans. \n\n")
	player.Position = 5
}

func fivef() {
	options := []string{"Look around.", "Search an exit", "Talk with the Snake"}

	if !player.GetKilled(0) {
		options = append(options, "Talk with the Stone")

	}
	answer := game.GetInputWithOptionsV(options, "What do you wanna do?")
	switch answer {
	case 0:
		player.Position = 6
		break
	case 1:
		player.Position = 7
		break
	case 2:
		game.WriteS("Snake: Go and look around.\n" +
			"Snake: But be careful!")
		break
	case 3:
		game.WriteS("Stone: That is so interesting ^-^\n" +
			"Stone: You are the first human i ever met.")
		break
	default:
		break
	}
}

func sixf() {
	game.WriteS("You look around.")
	game.WritelnSlow(". . .", 500)
	answer := rand.Intn(10)
	if answer < 4 {
		if LCgame.SetStage(GetFightStage(gobject.GetWolf())) == 1 {
			player.Kill(2)
			player.Position = 505
		} else {
			player.Position = 5
		}
	} else if answer < 9 && answer >= 4 {
		if LCgame.SetStage(GetFightStage(gobject.GetTiger())) == 1 {
			player.Position = 505
			player.Kill(3)
		} else {
			player.Position = 5
		}
	} else if answer == 9 {
		resFlowey := LCgame.SetStage(GetFightStage(gobject.GetFlowey()))
		if resFlowey == 2 {
			player.Position = 508
		} else if resFlowey == 1 {
			player.Position = 507
		} else if resFlowey == 0 {
			player.Position = 8
		}
	}
}

func sevenf() {
	game.WriteS("You go the way back.\n" +
		"You don't know exactly what you will do, but this was to crazy for you.\n" +
		"You see a flower at the ground.\n")
	resFlowey := LCgame.SetStage(GetFightStage(gobject.GetFlowey()))
	if resFlowey == 2 {
		player.Position = 508
	} else if resFlowey == 1 {
		player.Position = 507
	} else if resFlowey == 0 {
		player.Position = 8
	}
}

func eightf() {
	game.WriteS("Why does the flower know me...\n")
	game.WriteS("I think i should leave the forest.\n" +
		"I am scared.\n")
	player.Position = 100
	game.GetEnter()
}

func gtwof() {
	game.WriteS("Such an idiot.\n" +
		"After all i got, i know how to punish someone...\n\n" +
		"You go around and see a snake.\n")
	player.Position = 503
}

func gthreef() {
	if LCgame.SetStage(GetFightStage(gobject.GetSnake())) == 1 {
		player.Kill(1)
		player.Position = 504
	} else {
		player.Position = 4
	}
}

func gfourf() {
	game.WriteS("You have fun doing this.")
	player.Position = 505
}

func gfivef() {
	options := []string{"Look around for your next victim", "Search for an Exit", "Eat a piece of the Snake"}
	answer := game.GetInputWithOptionsV(options, "What do you wanna do?")
	switch answer {
	case 0:
		player.Position = 506
		break
	case 1:
		game.WriteS("You decided to leave this forest.\n" +
			"But you feel like somebody is following you.\n" +
			"You see a Flower\n")
		resFlowey := LCgame.SetStage(GetFightStage(gobject.GetFlowey()))
		if resFlowey == 2 {
			player.Position = 508
		} else if resFlowey == 1 {
			player.Position = 507
		} else if resFlowey == 0 {
			player.Position = 8
		}
		break
	case 2:
		game.WriteS("You recovered you hp")
		player.Chp = player.Hp
		break
	default:
		break
	}
}

func gsixf() {
	game.WriteS("You look around.")
	game.WriteSlow(". . .", 500)
	answer := rand.Intn(10)
	if answer < 4 {
		if LCgame.SetStage(GetFightStage(gobject.GetWolf())) == 1 {
			player.Kill(2)
			player.Position = 505
		} else {
			player.Position = 505
		}
	} else if answer < 9 && answer >= 4 {
		if LCgame.SetStage(GetFightStage(gobject.GetTiger())) == 1 {
			player.Position = 505
			player.Kill(3)
		} else {
			player.Position = 505
		}
	} else if answer == 9 {
		resFlowey := LCgame.SetStage(GetFightStage(gobject.GetFlowey()))
		if resFlowey == 2 {
			player.Position = 508
		} else if resFlowey == 1 {
			player.Position = 507
		} else if resFlowey == 0 {
			player.Position = 8
		}
	}
}

func gsevenf() {
	game.WriteS("What does the flower think who i am?\n" +
		"Now it's my time to rule this world.\n" +
		"*Laugh Laugh*\n")
	game.WritelnSlow("I will control everything.", 125)
	player.Kill(4)
	game.WriteS("You and Flowey are leaving the forest")
	player.Position = 509
	game.GetEnter()
}

func geightf() {
	game.WriteS("Flowey: We will fusion.\n" +
		"Flowey: We will be the strongest monster on earth\n" +
		"Flowey: Nobody will be able to beat us.\n\n")
	game.WriteS("You feel roots climbing your leg up.\n" +
		"You consist now of FLOWEY.")
	player.Name = "Fl" + strings.ToUpper(player.Name)
	for i := 0; i < 10; i++ {
		player.AddExp(gobject.GetFlowey())
	}
	if player.IsLevelUp() {
		game.WriteS(player.Name + " reached lv" + strconv.Itoa(player.Lv) + ".")
	}
	player.ConsistsFlowey = true
	player.Position = 509
}

func gninef() {
	if player.ConsistsFlowey {
		game.WriteS("You and Flowey are leaving the forest.")
	} else {
		game.WriteS("You are leaving the forest now.")
	}
	player.Position = 600
	game.GetEnter()
}
