package gobject

type Player struct {
	GObject
	Position          int
	ConsistsFlowey    bool
	SavedSoul         bool
	TalkedToDeadStone bool
	HitTheDoor        bool
	KilledEnemies     [21]bool
}

func NewPlayer(name string) Player {
	var pl Player
	pl.NewGObject(name, 10, 2, 3, 1, 1)
	return pl
}

func GetPlayer(pl Player) GObject {
	var lucy GObject
	lucy.NewGObject(pl.Name, pl.Hp*10, pl.Def*10, pl.Atk*10, pl.Atk*10, 1)
	lucy.StartSentence = "Even if you hacked the game.\nYou will have the same fate."
	lucy.DieSentence = "I am Lucy.\n" +
		"Also know as " + pl.Name
	lucy.Sentences = append(lucy.Sentences, "I will show you pain.",
		"I will show you yourself.",
		"I will show what you did.",
		"You feel scared.",
		"You realised you will die.")
	lucy.NextSentence = nextSentenceLucy
	return lucy
}

func nextSentenceLucy(g *GObject) string {
	g.Atk *= 10
	g.Def *= 10
	if g.PositionOfSentences < len(g.Sentences) {
		g.PositionOfSentences++
	}
	return g.Sentences[g.PositionOfSentences-1]
}

func (p Player) GetKilled(i int) bool {
	return p.KilledEnemies[i]
}

func (p *Player) Kill(i int) {
	p.KilledEnemies[i] = true
}
