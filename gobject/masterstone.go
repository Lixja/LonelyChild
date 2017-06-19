package gobject

func GetMasterStone() GObject {
	var mstone GObject
	mstone.NewGObject("MasterStone", 2, 1, 0, 1, 1)
	mstone.StartSentence = "A friendly Stone doesn't harm anybody."
	mstone.DieSentence = "Do you know what you just did?..."
	mstone.NeutralSentence = "The stone seems happy."
	mstone.Sentences = append(mstone.Sentences, "...")
	mstone.FightOptions = append(mstone.FightOptions, GetNewFightOption("Mercy", 1))
	return mstone
}
