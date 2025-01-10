package main

import "fmt"

type Enemy struct {
	pos    Vector2
	player *Player
	sprite string
	vel    Vector2
	damage uint
	health int
	id     int
}

// Simple movement towards player
func (e *Enemy) Step() {
	if e.player != nil {
		if e.player.pos.y > e.pos.y {
			e.vel.y++
		} else if e.player.pos.y < e.pos.y {
			e.vel.x--
		}

		if e.player.pos.x > e.pos.x {
			e.vel.x++
		} else if e.player.pos.x < e.pos.x {
			e.vel.x--
		}

		e.pos = addVector2(e.pos, e.vel)
	}

}

func (enemy Enemy) draw() {
	fmt.Printf("\033[s")
	fmt.Printf("\033[%d;%dH%s", enemy.pos.y, enemy.pos.x, enemy.sprite)
	fmt.Printf("\033[u")

}
