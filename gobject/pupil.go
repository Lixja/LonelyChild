package gobject

func GetPupil() GObject {
	var pupil GObject
	pupil.NewGObject("Pupil", 50, 10, 30, 11, 25)
	pupil.StartSentence = "An idiot. Show him what you can."
	pupil.DieSentence = "*SCREAMS*L.l..."
	pupil.NeutralSentence = "*SCREAMS*"
	pupil.Sentences = append(pupil.Sentences, "You are back.\nYou are really an idiot.")
	pupil.FightOptions = append(pupil.FightOptions, GetNewFightOption("Show him strength", 5),
		GetNewFightOption("Talk", 5))
	return pupil
}
