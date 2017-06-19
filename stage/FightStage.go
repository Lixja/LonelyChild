package stage

import (
	"LonelyChild/game"
	"LonelyChild/gobject"
	"math/rand"
	"os"
	"strconv"
)

func GetFightStage(enemy gobject.GObject) *game.Stage {
	fightStage := new(game.Stage)
	fightStage.Start = startFightStage
	fightStage.Enemy = enemy
	return fightStage
}

func startFightStage(s *game.Stage) {
	options := []string{"FIGHT", "ACT"}
	player := &s.LCgame.Data.Player
	enemy := s.Enemy
	game.WriteWall()
	game.WriteSlow("You will fight against: ", 80)
	game.WritelnSlow(s.Enemy.Name, 100)
	game.WriteWall()
	game.Writeln(enemy.StartSentence)
	game.WriteWall()
	for !enemy.IsDead() && !enemy.IsHelped() && !player.IsDead() {
		printStats(enemy)
		printStats(player.GObject)
		switch game.GetInputWithOptionsH(options, "Choice") {
		case 0:
			game.Writeln("You did " + strconv.Itoa(enemy.Damage(player.Atk)) + " damage to " + enemy.Name + ".")
			break
		case 1:
			i := game.GetInputWithOptionsH(enemy.GetFightOptionsAsString(), "Choice")
			enemy.Help(enemy.FightOptions[i].HelpValue)
			break
		}
		if !enemy.IsDead() && !enemy.IsHelped() {
			if !attackBlocked(enemy) {
				game.Writeln(enemy.Name + " attacked you!\nYou got " + strconv.Itoa(player.Damage(enemy.Atk)) + " damage.")
			}
			game.WriteWall()
			game.WriteS(enemy.NextSentence())
			game.WriteWall()
		}
	}
	if enemy.IsDead() {
		game.WritelnSlow(enemy.DieSentence, 50)
		game.Write("You got " + strconv.Itoa(player.AddExp(enemy)) + " exp")
		game.WritelnSlow("...", 50)
		if player.IsLevelUp() {
			game.Writeln("You are now level " + strconv.Itoa(player.Lv))
		}
	} else if enemy.IsHelped() {
		game.Writeln(enemy.NeutralSentence)
		player.Chp = player.Hp
	}
	game.WriteWall()
	if enemy.IsHelped() && !enemy.IsDead() {
		s.Happened = 0
	} else if !enemy.IsHelped() && enemy.IsDead() {
		s.Happened = 1
	} else if enemy.IsHelped() && enemy.IsDead() {
		s.Happened = 2
	} else if player.IsDead() {
		if enemy.Name != player.Name {
			game.WriteS("Y0u D1eD")
			os.Exit(0)
		}
	}
}

func attackBlocked(enemy gobject.GObject) bool {
	game.WriteS("Block the attack.")
	options := []string{"Stone", "Scissor", "Paper"}
	choice := game.GetInputWithOptionsH(options, "What will you choose?")
	cchoice := rand.Intn(3)
	game.WriteS(enemy.Name + " selected " + options[cchoice] + ".")
	if choice == cchoice {
		game.WriteS("Draw")
		return attackBlocked(enemy)
	} else if choice == cchoice-1 {
		game.WriteS("You Won")
		return true
	} else if choice == 2 && cchoice == 0 {
		game.WriteS("You Won")
		return true
	} else {
		game.WriteS("You Lost")
		return false
	}
}

func printStats(gob gobject.GObject) {
	game.Writeln(gob.Name + ": " + "HP: " + strconv.Itoa(gob.Hp) + "/" + strconv.Itoa(gob.Chp) + " ATK: " + strconv.Itoa(gob.Atk) + " DEF: " + strconv.Itoa(gob.Def))
}
