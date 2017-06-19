package gobject

func GetTeacher() GObject {
	var teacher GObject
	teacher.NewGObject("Teacher", 75, 25, 45, 12, 25)
	teacher.StartSentence = "The next idiot."
	teacher.DieSentence = "LU...."
	teacher.NeutralSentence = "*SCREAMS* LU..."
	teacher.Sentences = append(teacher.Sentences, "I will teach you how to respect me.")
	teacher.FightOptions = append(teacher.FightOptions, GetNewFightOption("Teach him", 5),
		GetNewFightOption("Talk", 10))
	return teacher
}
