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
	flowey.Help = help
	flowey.NextSentence = nextSentence
	return flowey
}

func help(g *GObject, i int) {
	if g.FightOptions[0].Option == "Accept" {
		if i == 0 {
			g.DieSentence = "Great decision partner."
			g.NeededHelp = 0
			g.Chp = 0
			g.Atk = 0
			g.Lv = 0
		} else if i == 1 {
			g.NeededHelp = 0
		}
	}
}

func nextSentence(g *GObject) string {
	if g.PositionOfSentences < len(g.Sentences) {
		g.PositionOfSentences++
		if g.PositionOfSentences == len(g.Sentences) {
			g.FightOptions = nil
			g.FightOptions = append(g.FightOptions, GetNewFightOption("Accept", 1), GetNewFightOption("Decline", 1))
		}
	}
	return g.Sentences[g.PositionOfSentences-1]
}
