package game

import (
	"time"
)

type Game struct {
	CurrentStage *Stage
	Data         GameData
}

type Stage struct {
	LCgame   *Game
	Start    func(*Stage)
	Happened int
}

func (g Game) Sleep(dur time.Duration) {
	time.Sleep(dur * time.Millisecond)
}

func (g Game) SetStage(stage *Stage) int {
	g.CurrentStage = stage
	g.CurrentStage.SetGame(g)
	g.CurrentStage.Start(g.CurrentStage)
	return g.CurrentStage.happen()
}

func (s Stage) SetGame(lcgame Game) {
	s.LCgame = &lcgame
}

func (g Game) ReadGameData() GameData {
	var gdata GameData
	g.Data = gdata
	return gdata
}

func (g Game) SaveGameData() {

}

func (s Stage) happen() int {
	return s.Happened
}
