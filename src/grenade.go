package main

import (
	"fmt"
)

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

func (g *Grenade) Step() {}

// TODO: Mke grenade explosion
func (g *Grenade) draw() {
	fmt.Printf("\033[s")
	sprite := g.trailSprite

	// TODO: Change to stepable code
	// LEGACY CODE
	fps := 10 // FPS here to make it run slower
	switch g.step {
	case 1 * fps:
		fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x) // All of these remove the old character
		g.pos.y -= 1 * g.amplitude
		g.pos.x++
		g.step++
	case 2 * fps:
		fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		g.pos.y -= 2 * g.amplitude
		g.pos.x++
		g.step++
	case 3 * fps:
		fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		g.pos.y += 2 * g.amplitude
		g.pos.x++
		g.step++
	case 4 * fps:
		fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		g.pos.y += 1 * g.amplitude
		g.pos.x++
		sprite = g.sprite
		g.step++
	case 5 * fps:
		// TODO: Make grenade explosion random
		// Just draws the thing
		// Empty: "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "
		new_sprite := "\0331A   !#\033[1B\033[4D$#@#$\033[1B\033[5D #@$%$#$#\033[1B\033[5D!@#"
		sprite = new_sprite
		g.step++
	case 6 * fps:
		new_sprite := "\0331A   ! \033[1B\033[4D @#@ \033[1B\033[5D  @#$%   \033[1B\033[5D * "
		sprite = new_sprite
		g.step++
	case 7 * fps:
		new_sprite := "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "
		sprite = new_sprite
		g.step++
	case 11 * fps:
		fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		gameManager.deleteObject(g.id, g.creationID) // kill self
	default:
		g.step++
	}

	fmt.Printf("\033[%d;%dH%s", g.pos.y, g.pos.x, sprite)
	fmt.Printf("\033[u")
}
