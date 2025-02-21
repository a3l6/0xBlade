package main

type Player struct {
	pos            Vector2
	lastPos        Vector2
	velocity       Vector2
	sprite         string
	keymap         Keymap
	last_direction uint8
	id             int
}

const DIRECTION_UP uint8 = 0
const DIRECTION_DOWN uint8 = 1
const DIRECTION_RIGHT uint8 = 2
const DIRECTION_LEFT uint8 = 3

func (p *Player) move(chars []uint8) {
	p.lastPos = p.pos
	for _, val := range chars {
		switch val {
		case p.keymap.up:
			p.last_direction = DIRECTION_UP
			p.velocity.y--
		case p.keymap.down:
			p.last_direction = DIRECTION_DOWN
			p.velocity.y++
		case p.keymap.left:
			p.last_direction = DIRECTION_LEFT
			p.velocity.x--
		case p.keymap.right:
			p.last_direction = DIRECTION_RIGHT
			p.velocity.x++
		}
	}
	newPos := addVector2(p.pos, p.velocity)
	CharOnNextPosition := gameManager.CurrBuffer[windowWidth*newPos.y+newPos.x]
	if CharOnNextPosition == 
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
	copy(gameManager.CurrBuffer[windowWidth*p.pos.y+p.pos.x:], []rune(p.sprite))
}
