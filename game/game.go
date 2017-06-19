package game

import (
	"LonelyChild/gobject"
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Game struct {
	CurrentStage *Stage
	Data         GameData
}

type Stage struct {
	LCgame   *Game
	Start    func(*Stage)
	Enemy    gobject.GObject
	Happened int
}

func NewGame() *Game {
	var LCgame Game
	LCgame.ReadGameData()
	return &LCgame
}

func Sleep(dur time.Duration) {
	time.Sleep(dur * time.Millisecond)
}

func (g Game) SetStage(stage *Stage) int {
	g.CurrentStage = stage
	g.CurrentStage.SetGame(&g)
	g.CurrentStage.Start(g.CurrentStage)
	return g.CurrentStage.happen()
}

func (s *Stage) SetGame(lcgame *Game) {
	s.LCgame = lcgame
}

func (g *Game) ReadGameData() {
	file, err := os.Open("gamedata.dat")
	if err != nil {
		g.Data.TeachedStory = false
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pos := 0
	for scanner.Scan() {
		switch pos {
		case 0:
			g.Data.TeachedStory = readEntryB(scanner.Text())
			break
		case 1:
			g.Data.Player = gobject.NewPlayer(scanner.Text())
			break
		case 2:
			g.Data.Player.Position, err = strconv.Atoi(scanner.Text())
			break
		case 3:
			g.Data.Player.Lv, err = strconv.Atoi(scanner.Text())
			break
		case 4:
			g.Data.Player.Atk, err = strconv.Atoi(scanner.Text())
			break
		case 5:
			g.Data.Player.Def, err = strconv.Atoi(scanner.Text())
			break
		case 6:
			g.Data.Player.Hp, err = strconv.Atoi(scanner.Text())
			break
		case 7:
			g.Data.Player.Chp, err = strconv.Atoi(scanner.Text())
			break
		case 8:
			g.Data.Player.ConsistsFlowey = readEntryB(scanner.Text())
			break
		case 9:
			g.Data.Player.SavedSoul = readEntryB(scanner.Text())
			break
		case 10:
			g.Data.Player.TalkedToDeadStone = readEntryB(scanner.Text())
			break
		case 11:
			g.Data.Player.HitTheDoor = readEntryB(scanner.Text())
			break
		}
		pos++
	}

}

func (g Game) SaveGameData() {
	var filec string
	data := g.Data
	if data.TeachedStory {
		filec += "1"
	} else {
		filec += "0"
	}
	filec += "\n"
	player := data.Player
	filec += player.Name
	filec += "\n"
	filec += createEntryI(player.Position)
	filec += createEntryI(player.Lv)
	filec += createEntryI(player.Atk)
	filec += createEntryI(player.Def)
	filec += createEntryI(player.Hp)
	filec += createEntryI(player.Chp)
	filec += createEntryB(player.ConsistsFlowey)
	filec += createEntryB(player.SavedSoul)
	filec += createEntryB(player.TalkedToDeadStone)
	filec += createEntryB(player.HitTheDoor)
	d := []byte(filec)
	ioutil.WriteFile("gamedata.dat", d, 0664)
}

func createEntryB(b bool) string {
	if b {
		return "1\n"
	} else {
		return "0\n"
	}
}

func readEntryB(b string) bool {
	if b == "1" {
		return true
	} else {
		return false
	}
}

func createEntryI(i int) string {
	return strconv.Itoa(i) + "\n"
}

func (s Stage) happen() int {
	return s.Happened
}
