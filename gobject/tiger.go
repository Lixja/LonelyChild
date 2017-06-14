package gobject

func GetTiger() GObject {
	var tiger GObject
	tiger.NewGObject("Tiger", 12, 1, 3, 4, 10)
	return tiger
}
