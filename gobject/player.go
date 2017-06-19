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

func (p Player) GetKilled(i int) bool {
	return p.KilledEnemies[i]
}

func (p *Player) kill(i int) {
	p.KilledEnemies[i] = true
}
