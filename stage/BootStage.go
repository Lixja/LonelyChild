package stage

import (
	"LonelyChild/game"
	"LonelyChild/gobject"
	"math/rand"
	"os"
	"strings"
)

func GetBootStage() *game.Stage {
	bootStage := new(game.Stage)
	bootStage.Start = start
	return bootStage
}

func start(s *game.Stage) {
	game.Writeln("___LonelyChild___")
	if s.LCgame.Data.TeachedStory {
		if s.LCgame.Data.Player.Position == 999 {
			game.WriteS("OnCe Up0n A TiMe...\n" +
				"ThEre WaS A LiTtlE Chi1d On A Ro11en W0rld.\n" +
				"Th1s W0r1d WaS OnCe A VeRy DaRk PlaCe.\n" +
				"But ThErE wAs A g0d nAmED.\n" +
				"LUCY\n")
			start := rand.Intn(3)
			if start == 2 {

			} else {
				//s.LCgame.setStage(new StoryStage());
			}
		} else if game.GetInputWithQuestionYesNo("Do you want to skip?") {
			//game.setStage(new StoryStage());
			return
		}
	}
	game.Sleep(1000)
	game.WriteS("Once Upon A Time...\n" +
		"There Was A Little Child On A Rotten World.\n" +
		"This World Was Once A Very Beautiful Place.\n" +
		"But The Human Race Destroyed This Place.\n" +
		"After Many Years A Child Was Born.\n" +
		"And The Humans Selected This Child To Be Their Next Victim.\n" +
		"The Story Of This Soul Is A Sad One.\n\n")
	game.WritelnSlow(". . .", 400)

	if s.LCgame.Data.Player.Position == 1234 {
		game.WriteS("This Child Run Away And Had A Big Journy.\n" +
			"After He Helped A Soul He Returned Home.\n" +
			"And Was Able To Come Together With His Family In Peace.\n" +
			"This Child Has Broken The Bad Side Of Humanity.\n" +
			"And Everybody Wondered Why They Were So Full Of Hate.\n\n")

		game.Sleep(500)
		game.WriteS("Th1s St0ry 1s S0 AkWaRd.\n")
		restart := game.GetInputWithQuestionYesNo("D0 y0u 1na rEsEt?")
		if restart {
			s.LCgame.Data.TeachedStory = false
			s.LCgame.Data.Player = gobject.NewPlayer("")
			s.LCgame.SaveGameData()
		} else {
			os.Exit(0)
		}
	} else if s.LCgame.Data.Player.Position == 1235 {
		game.WriteS("This Child Had Schizophrenia.\n" +
			"The Teacher Tried to Teach Him The Right Thing.\n" +
			"But He Refused.\n" +
			"The Pupils Were Bullied By Him.\n" +
			"And He Laugh.\n" +
			"His Parents Loved Him And Tried To Help.\n" +
			"But He Beat Them.\n" +
			"He Was A Bad Child.\n\n\n" +
			"Now The Police Found Corpse Of His Parents In His House.\n" +
			"The Child Was Just Sitting There And Laughing And Yelled A Name.\n\n" +
			"LUCY\n")
		restart := game.GetInputWithQuestionYesNo("D0 y0u 1na rEsEt?")
		if restart {
			s.LCgame.Data.TeachedStory = false
			s.LCgame.Data.Player = gobject.NewPlayer("")
			s.LCgame.SaveGameData()
		} else {
			os.Exit(0)
		}
	}
	game.WriteS("Every Human Bothered Him.\n" +
		"His Family Beat Him.\n" +
		"His \"Friends\" Bullied Him.\n" +
		"His Teacher Shouted At Him.\n" +
		"Even People He Did Not Knew Offended Him.\n\n" +
		"There Wasn't A Reason For Turtiring Him.\n" +
		"But Everybody Did It.\n" +
		"And He Could Not Understand Why Him.\n" +
		"So He Run Away.\n")

	game.WriteSlow("Far Far Away\n\n", 125)

	if !s.LCgame.Data.TeachedStory {
		name := selectName()
		s.LCgame.Data.Player = gobject.NewPlayer(name)
		s.LCgame.Data.TeachedStory = true
		s.LCgame.SaveGameData()
		game.Writeln("Right.\n" +
			"Now go and save your soul!")
		//game.setStage(new StoryStage());
		return
	} else {
		game.Writeln(s.LCgame.Data.Player.Name + "!!!")
		game.WriteS("Go and save yourself!")
		//game.setStage(new StoryStage());
	}
}

func selectName() string {
	name := game.GetInputWithQuestion("What was his name?")
	switch strings.ToLower(name) {
	case "lucy":
		game.WritelnSlow("This Name...", 250)
		game.WritelnSlow("HOW DO YOU KNOW THE NAME OF HELL!!!\n", 25)
		os.Exit(0)
		break
	case "flowey":
		game.WriteS("That is a flower not a name.")
		name = selectName()
		break
	case "fuck":
		game.WriteS("That is an insult not a name.")
		name = selectName()
		break
	case "josef":
		game.WriteS("That is the creators name not a name. ;P")
		name = selectName()
		break
	case "lonelychild":
		game.WriteS("Come on.\nBe creative.")
		name = selectName()
		break
	case "child":
		game.WriteS("Come on.\nBe creative.")
		name = selectName()
		break
	case "teacher":
		answer := game.GetInputWithQuestionYesNo("Are you a teacher?")
		if answer {
			game.WriteS("Be more creative.")
		} else {
			game.WriteS("Show the teacher that you can be more creative than him.")
		}
		name = selectName()
		break
	case "MasterStone":
		game.WriteS("YOU ARE NOT AN INNOCENT SOUL.")
		name = selectName()
		break
	}
	return name
}
