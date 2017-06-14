package gobject

func GetSnake() GObject {
	var snake GObject
	snake.NewGObject("Snake", 15, 1, 5, 2, 11)
	return snake
}
