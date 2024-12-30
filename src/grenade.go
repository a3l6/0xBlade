package main

import (
	"fmt"
)

func createNewGrenade(pos Vector2) Grenade {
	return Grenade{pos: pos, vel: Vector2{0, 0}, sprite: "O", trailSprite: "*", step: 0, amplitude: 1}
}

type Grenade struct {
	pos         Vector2
	vel         Vector2
	sprite      string
	trailSprite string
	step        int // between 1-4
	amplitude   int
	id          int
	creationID  uint8
}

// TODO: Mke grenade explosion
func (g *Grenade) draw() {
	fmt.Printf("\033[s")
	sprite := g.trailSprite
	switch g.step {
	case 0:
		g.step++
	case 1:
		g.pos.y -= 1 * g.amplitude
		g.pos.x++
		g.step++
	case 2:
		g.pos.y -= 2 * g.amplitude
		g.pos.x++
		g.step++
	case 3:
		g.pos.y += 2 * g.amplitude
		g.pos.x++
		g.step++
	case 4:
		g.pos.y += 1 * g.amplitude
		g.pos.x++
		sprite = g.sprite
		g.step++
	case 5:
		// TODO: Make grenade explosion random
		// Just draws the thing
		// Empty: "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "
		new_sprite := "\0331A   !#\033[1B\033[4D$#@#$\033[1B\033[5D #@$%$#$#\033[1B\033[5D!@#"
		sprite = new_sprite
		g.step++
	case 6:
		new_sprite := "\0331A   ! \033[1B\033[4D @#@ \033[1B\033[5D  @#$%   \033[1B\033[5D * "
		sprite = new_sprite
		g.step++
	case 7:
		new_sprite := "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "
		sprite = new_sprite
		g.step++
	default:
		gameManager.deleteObject(g.id, g.creationID) // kill self
	}
	fmt.Printf("\033[%d;%dH%s", g.pos.y, g.pos.x, sprite)
	fmt.Printf("\033[u")
}
