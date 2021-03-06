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

	//methods
	NextSentence func(g *GObject) string
	Help         func(g *GObject, i int)
	IsDead       func(g *GObject) bool
}

func (g *GObject) NewGObject(name string, hp, def, atk, lv, neededhelp int) {
	g.Name = name
	g.Hp = hp
	g.Chp = hp
	g.Def = def
	g.Atk = atk
	g.Lv = lv
	g.NeededHelp = neededhelp
	g.Exp = 0
	g.PositionOfSentences = 0
	g.NextSentence = NextSentence
	g.Help = Help
	g.IsDead = IsDead
}

func (g *GObject) Damage(atk int) int {
	damage := atk / g.Def
	g.Chp -= damage
	return damage
}

func (g *GObject) AddExp(gobj GObject) int {
	toaddexp := gobj.Lv * gobj.Hp
	g.Exp += toaddexp
	return toaddexp
}

func (g *GObject) IsLevelUp() bool {
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

func Help(g *GObject, v int) {
	g.NeededHelp -= v
}

func IsDead(g *GObject) bool {
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

func NextSentence(g *GObject) string {
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

func GetNewFightOption(option string, value int) FightOption {
	op := FightOption{Option: option, HelpValue: value}
	return op
}
