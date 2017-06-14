package gobject

func GetWolf() GObject {
	var wolf GObject
	wolf.NewGObject("Wolf", 10, 2, 4, 3, 15)
	return wolf
}
