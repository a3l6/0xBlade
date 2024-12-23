package main

import "fmt"

type Grenade struct {
	pos Vector2
	vel Vector2
	sprite string
	trailSprite string
	step uint // between 1-4
	amplitude int
	grenades *[]Grenade
}

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
	}
	fmt.Printf("\033[%d;%dH%s", g.pos.y, g.pos.x, sprite)
	fmt.Printf("\033[u")
}