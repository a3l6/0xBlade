package main

type Player struct {
	pos      Vector2
	lastPos  Vector2
	velocity Vector2
	sprite   string
	keymap   Keymap
	id       int
}

func (p *Player) move(chars []uint8) {
	p.lastPos = p.pos
	for _, val := range chars {
		switch val {
		case p.keymap.up:
			p.velocity.y--
		case p.keymap.down:
			p.velocity.y++
		case p.keymap.left:
			p.velocity.x--
		case p.keymap.right:
			p.velocity.x++
		}
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

func (p *Player) draw() {
	copy(gameManager.CurrBuffer[windowWidth*p.pos.y+p.pos.x:], []byte(p.sprite))
}
