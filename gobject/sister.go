package gobject

func GetSister(monster bool) GObject {
	var sister GObject
	sister.NewGObject("Sister", 100, 10, 30, 15, 30)
	sister.StartSentence = "Sister: You are not my Brother!"
	sister.DieSentence = "I hate you."
	sister.NeutralSentence = "ehm... i.. i am sorry."
	sister.Sentences = append(sister.Sentences, "Sister: Small little child.",
		"Sister: Couldn't my mother stop after our brother.",
		"Sister: Everybody hates you.",
		"Sister: I hate you.",
		"Sister: GO!")
	sister.FightOptions = append(sister.FightOptions, GetNewFightOption("Mercy", 1),
		GetNewFightOption("Talk", 5),
		GetNewFightOption("Forgive", 7))
	if monster {
		sister.Chp = 0
	}
	return sister
}
