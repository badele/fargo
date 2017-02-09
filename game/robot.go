package game

import (
	tl "github.com/badele/termloop"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Player structure
type Robot struct {
	*tl.Entity
	prevX int
	prevY int
	Level *tl.BaseLevel
}

// Draw implementation for Player
func (robot *Robot) Draw(screen *tl.Screen) {
	robot.Entity.Draw(screen)
}

func (robot *Robot) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		robot.prevX, robot.prevY = robot.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowLeft:
			robot.SetPosition(robot.prevX-1, robot.prevY)
			robot.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '◀'})
		case tl.KeyArrowRight:
			robot.SetPosition(robot.prevX+1, robot.prevY)
			robot.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▶'})
		case tl.KeyArrowUp:
			robot.SetPosition(robot.prevX, robot.prevY-1)
			robot.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▲'})
		case tl.KeyArrowDown:
			robot.SetPosition(robot.prevX, robot.prevY+1)
			robot.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '▼'})
		}
	}
}

func (robot *Robot) Collide(collision tl.Physical) {
	robot.SetPosition(robot.prevX, robot.prevY)
	// if _, ok := collision.(*tl.Entity); ok {
	// 	player.SetPosition(player.prevX, player.prevY)
	// }
}