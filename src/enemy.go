package main

import "fmt"

type Enemy struct {
	pos    Vector2
	sprite string
	vel    Vector2
	damage uint
	health int
}

func (self Enemy) draw() {
	fmt.Printf("\033[%d;%dH%s", self.pos.y, self.pos.x, self.sprite)
}
