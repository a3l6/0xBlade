package main

import "fmt"

type Enemy struct {
	pos    Vector2
	sprite string
	vel    Vector2
	damage uint
	health int
}

// Simple movement towards player
func (enemy *Enemy) step(playerPosition Vector2) {
	if playerPosition.x > enemy.pos.x {
		enemy.vel.x++
	} else if playerPosition.x < enemy.pos.x {
		enemy.vel.x--
	}

	if playerPosition.y > enemy.pos.y {
		enemy.vel.y++
	} else if playerPosition.y < enemy.pos.y {
		enemy.vel.x--
	}

	enemy.pos = addVector2(enemy.pos, enemy.vel)

}

func (enemy Enemy) draw() {
	fmt.Printf("\033[%d;%dH%s", enemy.pos.y, enemy.pos.x, enemy.sprite)
}
