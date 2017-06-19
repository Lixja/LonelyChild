package gobject

func GetSoul() GObject {
	var soul GObject
	soul.NewGObject("Soul", 5, 15, 10, 9, 100)
	soul.StartSentence = "Wow, i am really going crazy."
	soul.DieSentence = "The soul transforms into a deamon."
	soul.NeutralSentence = "The soul finds home."
	soul.Sentences = append(soul.Sentences, "hi?",
		"i am peter",
		"i was killed",
		"it was awful",
		"i can't rest",
		"i couldn't talk about that with somebody",
		"until now",
		"you seem friendly enough ^^",
		"thank you\ngood bye\n*smiles*")
	soul.FightOptions = append(soul.FightOptions, GetNewFightOption("listen", 10))
	return soul
}
