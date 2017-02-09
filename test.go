package main

import (
	"fmt"

	tl "github.com/badele/termloop"
	"github.com/badele/fargo/level"

)

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

func main() {
	game := tl.NewGame()
	mylevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	termwidth, termheight := game.InitialTermSize()
	array := level.BuildLevel(termwidth, termheight)
	level.ConverArrayToEntity(array, mylevel)
	level.InitBoard(game, mylevel)

	player := Player{
		Entity: tl.NewEntity(1, termheight-level.Chatheight-2, 1, 1),
		level:  mylevel,
	}

	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▶'})
	mylevel.AddEntity(&player)
	for i := 32; i < 512; i++ {
		fmt.Println(rune(i))
	}

	game.Screen().SetLevel(mylevel)
	game.Start()
}
