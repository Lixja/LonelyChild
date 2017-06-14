package gobject

func GetMasterStone() GObject {
	var mstone GObject
	mstone.NewGObject("MasterStone", 2, 1, 0, 1, 1)
	return mstone
}
