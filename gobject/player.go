package gobject

type Player struct {
	GObject
	Position          int
	ConsistsFlowey    bool
	SavedSoul         bool
	TalkedToDeadStone bool
	HitTheDoor        bool
}

func NewPlayer(name string) Player {
	var pl Player
	pl.NewGObject(name, 10, 2, 3, 1, 1)
	return pl
}
