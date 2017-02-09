package game

import (
	tl "github.com/badele/termloop"
	"github.com/badele/termloop/box"
)

const Chatheight int = 8
const Infowidth int = 30

func InitBoard(game *tl.Game, level *tl.BaseLevel) {
	//termx, termy := 80,25
	termwidth, termheight := game.InitialTermSize()

	// Right information block
	mybox := box.NewFrame(termwidth-Infowidth, 0, Infowidth, termheight-Chatheight, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Information ", box.AlignHCenter)
	mybox.LevelFollow(level)
	level.AddEntity(mybox)

	// Bottom chat panel
	mybox = box.NewFrame(0, termheight-Chatheight, termwidth, Chatheight, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Chat ", box.AlignHCenter)
	mybox.LevelFollow(level)
	level.AddEntity(mybox)

	// Bottom text area
	mytext := box.NewTextArea(1, termheight-Chatheight+1, termwidth-2, termheight-2, "", tl.ColorBlue, tl.ColorWhite, box.AlignNone)
	mytext.SetTypewriterDuration(100)
	mytext.SetText("This story takes place in the year 3000, You have a spacecraft for navigate", box.AlignNone)
	mytext.LevelFollow(level)
	level.AddEntity(mytext)
}

func BuildLevel(width, height int) [][]rune {
	//game *tl.Game, level *tl.BaseLevel

	h := height - Chatheight
	w := width - Infowidth
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

func ConverArrayToEntity(array [][]rune, level *tl.BaseLevel) {
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
