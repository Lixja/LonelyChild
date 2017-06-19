package gobject

func GetDeamon() GObject {
	var deamon GObject
	deamon.NewGObject("Deamon", 99, 30, 50, 10, 2)
	deamon.StartSentence = "You know you did a big mistake."
	deamon.DieSentence = "What are you...?\n" +
		"You killed me\n" +
		"I..I.. am sorry\n"
	deamon.NeutralSentence = "There is no neutral ending for this fight."
	deamon.Sentences = append(deamon.Sentences, "I will kill you.",
		"I will let you feel empty.",
		"You think it is just so simple to kill me!",
		"Feel my pain.")
	deamon.FightOptions = append(deamon.FightOptions, GetNewFightOption("DIE", 1))
	deamon.IsDead = isDead
	deamon.Help = helpDeamon
	return deamon
}

func isDead(g *GObject) bool {
	if g.NeededHelp == 2 {
		if g.Chp <= 0 {
			g.Chp = 66
			g.Sentences = append(g.Sentences,
				"Soul recovered hp from determination to kill you.",
				"I have to stay and kill you.")
		}
	} else {
		if g.Chp <= 0 {
			return true
		}
		return false
	}
	return false
}

func helpDeamon(g *GObject, v int) {
	g.Atk += v
}
