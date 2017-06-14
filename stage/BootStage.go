package stage

import "LonelyChild/game"

func GetBootStage() *game.Stage {
	bootStage := new(game.Stage)
	bootStage.Start = start
	bootStage.Happened = happened
	return bootStage
}

func start() {
	game.Writeln("___LonelyChild___")
	game.WriteS("Once Upon A Time...\n" +
		"There Was A Little Child On A Rotten World.\n" +
		"This World Was Once A Very Beautiful Place.\n" +
		"But The Human Race Destroyed This Place.\n" +
		"After Many Years A Child Was Born.\n" +
		"And The Humans Selected This Child To Be Their Next Victim.\n" +
		"The Story Of This Soul Is A Sad One.\n\n")
	game.WritelnSlow(". . .", 400)

}

func happened() int {
	return 0
}
