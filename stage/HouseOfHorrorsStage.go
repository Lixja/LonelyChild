package stage

import (
	"LonelyChild/game"
	"LonelyChild/gobject"
	"math/rand"
)

func GetHouseOfHorrorStage() *game.Stage {
	hohstage := new(game.Stage)
	hohstage.Start = startHOHStage
	return hohstage
}

func startHOHStage(s *game.Stage) {
	LCgame = s.LCgame
	player = &LCgame.Data.Player
	for (player.Position < 200 && player.Position >= 100) || (player.Position >= 600 && player.Position < 700) {
		switch player.Position {
		case 100:
			zeroh()
			break
		case 101:
			oneh()
			break
		case 102:
			twoh()
		case 103:
			threeh()
			break
		case 600:
			gzeroh()
			break
		case 601:
			goneh()
			break
		case 602:
			gtwoh()
			break
		case 603:
			gthreeh()
			break
		}
		LCgame.SaveGameData()
	}
}

func zeroh() {
	game.WriteS("*Walking ... Walking*\n" +
		"After several hours you found your way out of the forest.\n" +
		"You see a big house.\n" +
		"The door is open.\n" +
		"You enter...\n" +
		"The door closes immediatly\n" +
		"You can't leave.\n" +
		"You hear a voice.\n\n" +
		"Welcome to the house of horros.\n" +
		"You know you have to leave...\n\n")
	player.Position = 101
}

func oneh() {
	game.WriteWall()
	var options []string
	if player.SavedSoul {
		options = append(options, "Look around", "Call for help", "Hit the door")
	} else {
		options = append(options, "Look around", "Scream", "Hit the door")
	}
	answer := game.GetInputWithOptionsH(options, "What do you wanna do?")
	switch answer {
	case 0:
		player.Position = 102
		break
	case 1:
		if !player.SavedSoul {
			game.WriteS("You hear a scream from somebody else.")
			if LCgame.SetStage(GetFightStage(gobject.GetSoul())) == 1 {
				player.Kill(8)
				player.Position = (603)
			} else {
				player.SavedSoul = true
			}
		} else {
			player.Position = (103)
		}
		break
	case 2:
		game.WriteS("You hit the door\n" +
			". . .\n" +
			"Nothing happened.")
		break
	}
}

func twoh() {
	game.WriteS("You look around.")
	game.WriteSlow(". . .", 500)
	answer := rand.Intn(4)
	if answer == 0 {
		if LCgame.SetStage(GetFightStage(gobject.GetGhost())) == 1 {
			player.Kill(5)
			player.Position = (601)
		} else {
			player.Position = (101)
		}
	} else if answer == 1 {
		if LCgame.SetStage(GetFightStage(gobject.GetVampir())) == 1 {
			player.Position = (601)
			player.Kill(6)
		} else {
			player.Position = (101)
		}
	} else if answer == 2 {
		if LCgame.SetStage(GetFightStage(gobject.GetWereWolf())) == 1 {
			player.Kill(7)
			player.Position = (601)
		} else {
			player.Position = (101)
		}
	}
}

func threeh() {
	game.WriteS("A soul appeared to help you.\n" +
		"You close your eyes.\n" +
		"You open them again.\n" +
		"You see\n" +
		"You see your ... home.\n\n" +
		"Why am i back at this place you ask.\n" +
		"Soul: Believe me, this time things will change.")
	player.Position = (200)
}

func gzeroh() {
	game.WriteS("*Walking ... Walking*\n")
	if player.ConsistsFlowey {
		game.WriteS("After several hours you and flowey found your way out of the forest.\n")
	} else {
		game.WriteS("After several hours you found your way out of the forest.\n")
	}
	game.WriteS("You see a big house.\n" +
		"The door is open.\n" +
		"You enter...\n" +
		"The door closes immediatly\n" +
		"You can't leave.\n" +
		"You hear a voice.\n\n" +
		"Welcome dear friend.\n" +
		"This is your new home lost soul.\n")
	player.Position = (601)
}

func goneh() {
	game.WriteWall()
	var options []string
	if player.ConsistsFlowey {
		options = append(options, "Look around", "Hit the door", "Destroy the door")
	} else {
		options = append(options, "Look around", "Hit the door")
	}
	if player.HitTheDoor {
		options[1] = "Talk to the soul"
	}
	if player.GetKilled(0) && !player.TalkedToDeadStone {
		options = append(options, "Talk to Stone.")
	}
	answer := game.GetInputWithOptionsH(options, "What do you wanna do?")
	switch answer {
	case 0:
		player.Position = 602
		break
	case 1:
		if player.HitTheDoor {
			game.WriteS("Soul: Will you free me?\n")
			if LCgame.SetStage(GetFightStage(gobject.GetSoul())) == 1 {
				player.Kill(8)
				player.Position = 603
			} else {
				player.Position = 200
			}
		} else {
			game.WriteS("You hit the door\n" +
				". . .\n" +
				"Nothing happened except that somebody woke up.")
			player.HitTheDoor = true
		}
		break
	case 2:
		if options[2] != "Talk to Stone" {
			game.WriteS("You and flowey destroy the door.\n" +
				"Something appears in front of you.\n")
			if LCgame.SetStage(GetFightStage(gobject.GetMystery())) == 1 {
				player.Kill(10)
				player.Position = 601
				game.WriteS("Mystery built a new door.")
			} else {
				player.Position = 601
			}
			break
		}
	case 3:
		game.WriteS("Stone: Hello...\n" +
			"Stone: Do you remember when you killed me?\n" +
			"Stone: But this is you last chance.\n" +
			"Stone: Do the right thing!\n")
		player.TalkedToDeadStone = true
		break
	}
}

func gtwoh() {
	game.WriteS("You look around.")
	game.WriteSlow(". . .", 500)
	answer := rand.Intn(4)
	if answer == 0 {
		if LCgame.SetStage(GetFightStage(gobject.GetGhost())) == 1 {
			player.Kill(5)
			player.Position = (601)
		} else {
			player.Position = (601)
		}
	} else if answer == 1 {
		if LCgame.SetStage(GetFightStage(gobject.GetVampir())) == 1 {
			player.Position = (601)
			player.Kill(6)
		} else {
			player.Position = (601)
		}
	} else if answer == 2 {
		if LCgame.SetStage(GetFightStage(gobject.GetWereWolf())) == 1 {
			player.Kill(7)
			player.Position = (601)
		} else {
			player.Position = (601)
		}
	} else if answer == 3 {
		if LCgame.SetStage(GetFightStage(gobject.GetSoul())) == 1 {
			player.Kill(8)
			player.Position = (603)
		} else {
			player.Position = (200)
		}
	}
}

func gthreeh() {
	if LCgame.SetStage(GetFightStage(gobject.GetDeamon())) == 1 {
		player.Kill(9)
		player.Position = (604)
	}
}

func gfourh() {
	game.WriteS("The house starts to burn.\n" +
		"Mystery opens the door.\n" +
		"You left the house.\n" +
		"After some minutes there is no house at this place any more.\n" +
		"There is only a deamon\n\n")
	game.WritelnSlow("YOU", 200)
	player.Position = (700)
}
