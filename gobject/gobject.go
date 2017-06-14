package gobject

import "math/rand"

type FightOption struct {
	Option    string
	HelpValue int
}

type GObject struct {
	Name                string
	Hp                  int
	Chp                 int
	Def                 int
	Atk                 int
	Lv                  int
	Exp                 int
	NeededHelp          int
	Seededhelp          int
	Sentences           []string
	PositionOfSentences int
	FightOptions        []FightOption
	StartSentence       string
	DieSentence         string
	NeutralSentence     string
}

func (g GObject) Damage(atk int) int {
	damage := atk / g.Def
	g.Chp -= damage
	return damage
}

func (g GObject) AddExp(gobj GObject) int {
	toaddexp := gobj.Lv * gobj.Hp
	g.Exp += toaddexp
	return toaddexp
}

func (g GObject) IsLevelUp() bool {
	if g.Exp > (g.Lv * g.Atk) {
		g.Lv++
		g.Hp += rand.Int()*g.Lv + 1
		g.Chp = g.Hp
		g.Atk += rand.Int()*g.Lv + 1
		g.Def += rand.Int()*g.Lv + 1
		g.Exp = 0
		return true
	}
	return false
}

func (g GObject) Help(v int) {
	g.NeededHelp -= v
}

func (g GObject) IsDead() bool {
	if g.Chp <= 0 {
		return true
	}
	return false
}

func (g GObject) IsHelped() bool {
	if g.NeededHelp <= 0 {
		return true
	}
	return false
}

func (g GObject) NextSentence() string {
	if g.PositionOfSentences < len(g.Sentences) {
		g.PositionOfSentences++
	}
	return g.Sentences[g.PositionOfSentences-1]
}

func (g GObject) GetFightOptionsAsString() []string {
	var foptions []string
	for _, f := range g.FightOptions {
		foptions = append(foptions, f.Option)
	}
	return foptions
}
