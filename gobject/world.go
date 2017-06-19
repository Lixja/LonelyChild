package gobject

func GetWorld() GObject {
	var world GObject
	world.NewGObject("World", 500, 100, 150, 18, 1)
	world.StartSentence = "Are you sure it's a great idea to destroy the world?\n" +
		"I mean you live there.\n" +
		"P.s. and this pc is also there.\n"
	world.DieSentence = "THVjeQ: You did such a big mistake.\n" +
		"THVjeQ: Now suffer.\n"
	world.NeutralSentence = ""
	world.Sentences = append(world.Sentences, "Do you really think you can destroy me that easily?",
		"You know there are consequences, am I right?",
		"What would you do if ... L.Lucy would appear?",
		"I a planet.\n",
		"You are a demon.",
		"...")
	return world
}
