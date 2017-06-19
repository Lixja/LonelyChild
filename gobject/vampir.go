package gobject

func GetVampir() GObject {
	var vampir GObject
	vampir.NewGObject("Vampir", 50, 1, 10, 7, 15)
	vampir.StartSentence = "Not a human.\nOh it's a vampir."
	vampir.DieSentence = "*Scream and decays into ash."
	vampir.NeutralSentence = "Vampir: You was delicious."
	vampir.Sentences = append(vampir.Sentences, "I love human blood.",
		"I would do anything for blood.")
	vampir.FightOptions = append(vampir.FightOptions, GetNewFightOption("Spend Blood", 15))
	return vampir
}
