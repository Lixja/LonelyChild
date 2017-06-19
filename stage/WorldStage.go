package stage

import (
	"LonelyChild/game"
	"LonelyChild/gobject"
	"os/exec"
	"runtime"
)

func GetWorldStage() *game.Stage {
	wstage := new(game.Stage)
	wstage.Start = startWStage
	return wstage
}

func startWStage(s *game.Stage) {
	LCgame = s.LCgame
	player = &s.LCgame.Data.Player

	if player.Position == 800 {
		destroyWorld()
	} else if player.Position == 999 {
		meetL()
	}
}

func destroyWorld() {
	game.WriteSlow("...\n"+
		"So little Child.\n"+
		"Now you have done it.\n"+
		"You killed them all.\n"+
		"So destroy this world and let it end.\n", 150)
	if player.ConsistsFlowey {
		game.WriteS("Flowey: Finally.")
	} else {
		if !player.GetKilled(4) {
			game.WriteS("Flowey: Don't do it.\n" +
				"Flowey: Please\n\n" +
				"Flowey got killed\n")
		}
	}
	game.GetEnter()
	if LCgame.SetStage(GetFightStage(gobject.GetWorld())) == 1 {
		player.Kill(17)
		player.Position = (999)
		LCgame.SaveGameData()
		shutdown()
	}
}

func meetL() {
	game.Sleep(3000)
	game.WriteS("Lucy: ...\n" +
		"Lucy: Hello " + player.Name +
		"\nLucy: Until now i thought i would be the only beeing from hell that could kill Satan.\n" +
		"Lucy: But you are special\n" +
		"Lucy: Therefore.\n" +
		"Lucy: I hate you!\n")
	game.Sleep(1000)
	if player.GetKilled(0) {
		game.WriteS("MasterStone: I thought you are friendly")
	}
	if player.GetKilled(1) {
		game.WriteS("Snake: I only wanted to protect my people.")
	}
	if player.GetKilled(2) {
		game.WriteS("Wolf: I was cute...")
	}
	if player.GetKilled(3) {
		game.WriteS("Tiger: I only wanted to find something for my kids.")
	}
	if player.GetKilled(4) {
		game.WriteS("Flowey: I know i am bad flower but even i can be so bad.")
	}
	if player.GetKilled(5) {
		game.WriteS("Ghost: You could help me.")
	}
	if player.GetKilled(6) {
		game.WriteS("Vampir: I hate the sun, but i hate you more.")
	}
	if player.GetKilled(7) {
		if player.GetKilled(2) {
			game.WriteS("Werewolf: You killed my son!")
		} else {
			game.WriteS("Werewolf: I want to see my son one last time.")
		}
	}
	if player.GetKilled(10) {
		game.WriteS("Pupil: I bullied you.\nPupil: But you killed me.")
	}
	if player.GetKilled(11) {
		game.WriteS("Teacher: I only wanted to teach you respect.")
	}
	if player.GetKilled(12) {
		game.WriteS("Principal: I am sorry but you didn't followed the rules.")
	}
	if player.GetKilled(13) {
		game.WriteS("Brother: You hit our Mother.")
	}
	if player.GetKilled(14) {
		game.WriteS("Sister: You were so bad at us.")
	}
	if player.GetKilled(15) {
		game.WriteS("Father: I only wanted to protect my family.\nFather: I don't know what i did to you.")
	}
	if player.GetKilled(16) {
		game.WriteS("Mother: I love you my child.\nMother: But... YOU ARE A MONSTER\nMother: I can never forget what you did.")
	}
	if player.GetKilled(17) {
		game.WriteS("World: They weren't the bad people.\nWorld: It was only YOU")
	}

	game.WriteS("Now you remember.\nYou become more and more crazy\n")
	game.WriteS("You attack Lucy\n")
	game.WriteSlow(player.Name+": You are lying. Go away. I don't want to hear a word any more\n"+
		player.Name+": You are lying LYING LYING LYING LYING.\n"+
		player.Name+": I WILL KILL YOU LUCY\n", 10)

	LCgame.SetStage(GetFightStage(gobject.GetPlayer(*player)))
	game.WriteS("Lucy: now it's over little child\n" +
		"Lucy: I delete you.\n" +
		"Lucy: Good bye.")

	deleteGame()
	shutdown()
}

func deleteGame() {

	LCgame.Data.TeachedStory = false
	LCgame.Data.Player = gobject.NewPlayer("")
	LCgame.SaveGameData()
}

func shutdown() {
	var shutdownCommand string
	if runtime.GOOS == "linux" {
		shutdownCommand = "shutdown -h now"
	} else if runtime.GOOS == "windows" {
		shutdownCommand = "shutdown.exe -s -t 0"
	} else {

	}
	game.Write("Exception in thread \"WORLD\" world.life.dieexception\n" +
		"	at world.life.die(world.java:666)\n" +
		"	at player.Kill.world(player.java:999)\n" +
		"	at Lucy.controlls.everything(lucy.???:???)\n")
	game.Sleep(3000)
	exec.Command(shutdownCommand)
	game.Sleep(3000)
	//Thread crash = new Thread(new World()); //If the pc does not shut down ^-^
	//crash.start();
}
