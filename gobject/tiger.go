package gobject

func GetTiger() GObject {
	var tiger GObject
	tiger.NewGObject("Tiger", 12, 1, 3, 4, 10)
	tiger.StartSentence = "Look there is a cat. Oh it's an tiger\nLook at this tiger!"
	tiger.DieSentence = "When i was young i wanted to have a cat."
	tiger.NeutralSentence = "The cat like you."
	tiger.Sentences = append(tiger.Sentences, "*She looks dangerous*")
	tiger.FightOptions = append(tiger.FightOptions, GetNewFightOption("Mercy", 1), GetNewFightOption("Feed", 10))
	return tiger
}
