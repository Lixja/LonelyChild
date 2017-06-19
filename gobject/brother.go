package gobject

func GetBrother(monster bool) GObject {
	var brother GObject
	brother.NewGObject("Brother", 100, 30, 50, 14, 30)
	brother.StartSentence = "Brother: You are not my Brother!"
	brother.DieSentence = "I hate you."
	brother.NeutralSentence = "ehm... i.. i am sorry."
	brother.Sentences = append(brother.Sentences, "Brother: Small little child.",
		"Brother: Couldn't my mother stop after me.",
		"Brother: Everybody hates you.",
		"Brother: Why did you came back?",
		"Brother: GO!")
	brother.FightOptions = append(brother.FightOptions, GetNewFightOption("Mercy", 1),
		GetNewFightOption("Talk", 5),
		GetNewFightOption("Forgive", 7))
	if monster {
		brother.Chp = 0
	}
	return brother
}
