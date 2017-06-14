package gobject

func GetFlowey() GObject {
	var flowey GObject
	flowey.NewGObject("Flowey", 20, 1, 4, 5, 1)
	return flowey
}
