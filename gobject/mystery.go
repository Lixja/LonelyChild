package gobject

func GetMystery() GObject {
	var mystery GObject
	mystery.NewGObject("Mystery", 100, 50, 50, 10, 9)
	mystery.StartSentence = "Ehmm, you don't know what is infront of you."
	mystery.DieSentence = "R28gaG9tZSBhbmQgZGllLg"
	mystery.NeutralSentence = "Find your way home!"
	mystery.Sentences = append(mystery.Sentences, "???")
	mystery.FightOptions = append(mystery.FightOptions, GetNewFightOption("???", 0),
		GetNewFightOption("???", 0),
		GetNewFightOption("???", 0))
	return mystery
}
