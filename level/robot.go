package level

import (
	tl "github.com/badele/termloop"
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