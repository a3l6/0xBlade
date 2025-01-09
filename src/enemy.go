package main

import "fmt"

type Enemy struct {
	pos       Vector2
	playerPos *Vector2
	sprite    string
	vel       Vector2
	damage    uint
	health    int
	id        int
}

// Simple movement towards player
func (enemy *Enemy) Step() {
	if (*enemy.playerPos).x > enemy.pos.x {
		enemy.vel.x++
	} else if (*enemy.playerPos).x < enemy.pos.x {
		enemy.vel.x--
	}

	if (*enemy.playerPos).y > enemy.pos.y {
		enemy.vel.y++
	} else if (*enemy.playerPos).y < enemy.pos.y {
		enemy.vel.x--
	}

	enemy.pos = addVector2(enemy.pos, enemy.vel)

}

func (enemy Enemy) draw() {
	fmt.Printf("\033[s")
	fmt.Printf("\033[%d;%dH%s", enemy.pos.y, enemy.pos.x, enemy.sprite)
	fmt.Printf("\033[u")

}
