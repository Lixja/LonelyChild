package gobject

type Player struct {
	GObject
	Position          int
	ConsistsFlowey    bool
	SavedSoul         bool
	TalkedToDeadStone bool
	HitTheDoor        bool
}
