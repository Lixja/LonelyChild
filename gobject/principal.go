package gobject

func GetPrincipal() GObject {
	var principal GObject
	principal.NewGObject("Principal", 100, 25, 50, 13, 25)
	principal.StartSentence = "The last idiot from this place."
	principal.DieSentence = "no..NO.NOOOOOOO"
	principal.NeutralSentence = "*SCREAMS*"
	principal.Sentences = append(principal.Sentences, "You will get suspended.")
	principal.FightOptions = append(principal.FightOptions, GetNewFightOption("Show respect", 5),
		GetNewFightOption("Talk", 5))
	return principal
}
