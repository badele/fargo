package main

import (
	"fmt"

	tl "github.com/badele/termloop"
	"github.com/badele/termloop/box"
)

const chatheight int = 8
const infowidth int = 30

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Player structure
type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

// Draw implementation for Player
func (player *Player) Draw(screen *tl.Screen) {
	player.Entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
			player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '◀'})
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
			player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▶'})
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
			player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▲'})
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
			player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▼'})
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	player.SetPosition(player.prevX, player.prevY)
	// if _, ok := collision.(*tl.Entity); ok {
	// 	player.SetPosition(player.prevX, player.prevY)
	// }
}

func initBoard(game *tl.Game, level *tl.BaseLevel) {
	//termx, termy := 80,25
	termwidth, termheight := game.InitialTermSize()

	// Right information block
	mybox := box.NewFrame(termwidth-infowidth, 0, infowidth, termheight-chatheight, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Information ", box.AlignHCenter)
	mybox.LevelFollow(level)
	level.AddEntity(mybox)

	// Bottom chat panel
	mybox = box.NewFrame(0, termheight-chatheight, termwidth, chatheight, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Chat ", box.AlignHCenter)
	mybox.LevelFollow(level)
	level.AddEntity(mybox)

	// Bottom text area
	mytext := box.NewTextArea(1, termheight-chatheight+1, termwidth-2, termheight-2, "", tl.ColorBlue, tl.ColorWhite, box.AlignNone)
	mytext.SetTypewriterDuration(100)
	mytext.SetText("This story takes place in the year 3000, You have a spacecraft for navigate", box.AlignNone)
	mytext.LevelFollow(level)
	level.AddEntity(mytext)
}

func buildLevel(width, height int) [][]rune {
	//game *tl.Game, level *tl.BaseLevel

	h := height - chatheight
	w := width - infowidth
	middlex := w / 2
	middley := h / 2

	array := make([][]rune, h)
	for y := 0; y < h; y++ {
		array[y] = make([]rune, w)
		for x := 0; x < w; x++ {
			array[y][x] = ' '
		}
	}

	// Horizontal line
	for x := 0; x < w; x++ {
		array[0][x] = '█'
		array[h-1][x] = '█'
	}

	// Vertical line
	for y := 0; y < h; y++ {
		array[y][0] = '█'
		array[y][w-1] = '█'
	}

	// Cardinal point
	array[1][middlex] = 'N'
	array[h-2][middlex] = 'S'
	array[middley][1] = 'W'
	array[middley][w-2] = 'E'

	return array
}

func converArrayToEntity(array [][]rune, level *tl.BaseLevel) {
	lenx := 0
	leny := len(array)

	if leny > 0 {
		lenx = len(array[0])
	}

	for y := 0; y < leny; y++ {
		for x := 0; x < lenx; x++ {
			ch := array[y][x]
			if ch != ' ' {
				obj := tl.NewText(x, y, string(ch), tl.ColorWhite, tl.ColorBlack)
				level.AddEntity(obj)
			}
		}
	}
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	termwidth, termheight := game.InitialTermSize()
	array := buildLevel(termwidth, termheight)
	converArrayToEntity(array, level)
	initBoard(game, level)

	player := Player{
		Entity: tl.NewEntity(1, termheight-chatheight-2, 1, 1),
		level:  level,
	}

	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▶'})
	level.AddEntity(&player)
	for i := 32; i < 512; i++ {
		fmt.Println(rune(i))
	}

	game.Screen().SetLevel(level)
	game.Start()
}
