package gobject

func GetGhost() GObject {
	var ghost GObject
	ghost.NewGObject("Ghost", 30, 10, 10, 6, 15)
	ghost.StartSentence = "A Ghost. Buh."
	ghost.DieSentence = "Ghosts can't die ;P\nBut I hate you."
	ghost.NeutralSentence = "The ghost finds peace."
	ghost.Sentences = append(ghost.Sentences, "buuuuh.")
	ghost.FightOptions = append(ghost.FightOptions,
		GetNewFightOption("Mercy", 1),
		GetNewFightOption("Talk", 6),
		GetNewFightOption("Say Buh", 6))
	return ghost
}
