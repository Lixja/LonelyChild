package gobject

func GetWolf() GObject {
	var wolf GObject
	wolf.NewGObject("Wolf", 10, 2, 4, 3, 15)
	wolf.StartSentence = "A Wolf, he looks cute and hungry."
	wolf.DieSentence = "He looks still cute, even if you see his intestines."
	wolf.NeutralSentence = "He still wants to be pet."
	wolf.Sentences = append(wolf.Sentences, "*heckle*")
	wolf.FightOptions = append(wolf.FightOptions, GetNewFightOption("Mercy", 1), GetNewFightOption("Pet", 10), GetNewFightOption("Growl", -2))
	return wolf
}
