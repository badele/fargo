package main

import (
	"github.com/badele/fargo/game"
	tl "github.com/badele/termloop"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	mygame := tl.NewGame()
	mylevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	termwidth, termheight := mygame.InitialTermSize()
	array := game.BuildLevel(termwidth, termheight)
	game.ConverArrayToEntity(array, mylevel)
	game.InitBoard(mygame, mylevel)

	robot := game.Robot{
		Entity: tl.NewEntity(1, termheight-game.Chatheight-2, 1, 1),
		Level:  mylevel,
	}

	// Set the character at position (0, 0) on the entity.
	robot.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▶'})
	mylevel.AddEntity(&robot)

	mygame.Screen().SetLevel(mylevel)
	mygame.Start()
}
