package gobject

func GetSnake() GObject {
	var snake GObject
	snake.NewGObject("Snake", 15, 1, 5, 2, 11)
	snake.StartSentence = "A Snake, An angry Snake."
	snake.DieSentence = "The Snake tries to rest in peace.\n Do you have fun?"
	snake.NeutralSentence = "The Snakes apologizes to you."
	snake.Sentences = append(snake.Sentences, "The Snake seems to wanna hurt you!")
	snake.FightOptions = append(snake.FightOptions, GetNewFightOption("Mercy", 1), GetNewFightOption("Talk", 10))
	return snake
}
