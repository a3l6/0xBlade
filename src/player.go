package main

import "fmt"

type Player struct {
	pos      Vector2
	velocity Vector2
	sprite   string
	l        *Level
	keymap   Keymap
}

func (p *Player) move(char uint8) {
	switch char {
	case p.keymap.up:
		fmt.Printf("\033[%d;%dH ", p.pos.y, p.pos.x)
		p.velocity.y--
	case p.keymap.down:
		fmt.Printf("\033[%d;%dH ", p.pos.y, p.pos.x)
		p.velocity.y++
	case p.keymap.left:
		fmt.Printf("\033[%d;%dH ", p.pos.y, p.pos.x)
		p.velocity.x--
	case p.keymap.right:
		fmt.Printf("\033[%d;%dH ", p.pos.y, p.pos.x)
		p.velocity.x++
	}

	newPos := addVector2(p.pos, p.velocity)
	if newPos.x < (*p.l).leftBound {
		newPos.x++
	}
	if newPos.x > (*p.l).rightBound {
		newPos.x--
	}
	if newPos.y < (*p.l).upperBound {
		newPos.y++
	}
	if newPos.y > (*p.l).lowerBound {
		newPos.y--
	}

	p.pos = newPos
	p.velocity = Vector2{x: 0, y: 0}

}

func (p *Player) draw() {
	fmt.Printf("\033[%d;%dH%s", p.pos.y, p.pos.x, p.sprite)
}
