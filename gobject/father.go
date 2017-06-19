package gobject

func GetFather(monster bool) GObject {
	var father GObject
	father.NewGObject("Father", 150, 50, 50, 16, 30)
	father.StartSentence = "Father: I am disapointed from you."
	father.DieSentence = "I am still disapointed from you."
	father.NeutralSentence = "My son."
	father.Sentences = append(father.Sentences, "Father: Why did i have to bang your mother without the condom.",
		"Father: You can't understand how much i hate you.",
		"Father: I will never love you",
		"Father: go away.")
	father.FightOptions = append(father.FightOptions,
		GetNewFightOption("Mercy", 1),
		GetNewFightOption("Talk", 5),
		GetNewFightOption("Forgive", 7),
		GetNewFightOption("Remember i am your child", 8))
	if monster {
		father.Chp = 0
	}
	return father
}
