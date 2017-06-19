package gobject

func GetFlowey() GObject {
	var flowey GObject
	flowey.NewGObject("Flowey", 20, 1, 4, 5, 1)
	flowey.StartSentence = "A Flower?!"
	flowey.DieSentence = "L..U..C...Y"
	flowey.NeutralSentence = "..I will still follow you little guy."
	flowey.Sentences = append(flowey.Sentences, "I know you.", "I know you better than you think.", "I saw how you was tortured",
		"I know what you want.", "Should I help you?")
	flowey.FightOptions = append(flowey.FightOptions, GetNewFightOption("Mercy", 0), GetNewFightOption("Talk", 0))
	return flowey
}
