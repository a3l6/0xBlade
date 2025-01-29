package main

type Player struct {
	pos      Vector2
	velocity Vector2
	sprite   rune
	id       int
}

func (p *Player) Move(buf []byte) {
	switch buf[0] {
	case keymap.up:
		p.velocity.y--
	case keymap.down:
		p.velocity.y++
	case keymap.left:
		p.velocity.x--
	case keymap.right:
		p.velocity.x++
	}

	newPos := addVector2(p.pos, p.velocity)
	if newPos.x < level.leftBound {
		newPos.x++
	}
	if newPos.x > level.rightBound {
		newPos.x--
	}
	if newPos.y < level.upperBound {
		newPos.y++
	}
	if newPos.y > level.lowerBound {
		newPos.y--
	}

	p.pos = newPos
	p.velocity = Vector2{x: 0, y: 0}
}

func (p *Player) collision() {}

func (p *Player) Draw() {
	level.print(p.pos, LevelChar{master: p.id, char: p.sprite})
}
