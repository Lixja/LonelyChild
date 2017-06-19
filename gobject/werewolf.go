package gobject

func GetWereWolf() GObject {
	var wwolf GObject
	wwolf.NewGObject("Werewolf", 50, 10, 1, 8, 15)
	wwolf.StartSentence = "Uhhh, a Werewolf ^-^."
	wwolf.DieSentence = "He looks like his father from the forest."
	wwolf.NeutralSentence = "He looks kind of cute."
	wwolf.Sentences = append(wwolf.Sentences, "I am a wolf, what should i say?")
	wwolf.FightOptions = append(wwolf.FightOptions, GetNewFightOption("pet", 3),
		GetNewFightOption("Pet", 6),
		GetNewFightOption("PET", 9))
	return wwolf
}
