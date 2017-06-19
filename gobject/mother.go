package gobject

func GetMother(monster bool) GObject {
	var mother GObject
	mother.NewGObject("Mother", 150, 50, 60, 17, 30)
	mother.StartSentence = "You."
	mother.DieSentence = "I did never love you."
	mother.NeutralSentence = "I love you.\nI am so sorry that i was such a bad mother.\n"
	mother.Sentences = append(mother.Sentences, "Mother: Why do you have to be my son.",
		"Mother: Why can't you be like your brother or sister.",
		"Mother: You are the dead for us.",
		"Mother: Hurry Up and GO AWAY!",
		"Mother: GO!")
	mother.FightOptions = append(mother.FightOptions,
		GetNewFightOption("Mercy", 1),
		GetNewFightOption("Talk", 5),
		GetNewFightOption("Forgive", 7),
		GetNewFightOption("Remember i am your child", 8))
	if monster {
		mother.Chp = 0
	}
	return mother
}
