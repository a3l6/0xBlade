package main

type Enemy struct {
	pos        fVector2
	lastPos    fVector2
	sprite     rune
	vel        Vector2
	damage     uint
	health     int
	id         int
	creationId int
}

// Simple movement towards player
func (e *Enemy) Step() {
	if gameManager.ptrPlayer != nil {
		const speed = .5

		e.lastPos = e.pos
		if gameManager.ptrPlayer.pos.x > int(e.pos.x) {

			charOnNextPosition := gameManager.CurrBuffer[windowWidth*int(e.pos.y)+int(e.pos.x)+1]
			if charOnNextPosition != e.sprite {
				e.pos.x += speed
			}

		} else if gameManager.ptrPlayer.pos.x < int(e.pos.x) {

			charOnNextPosition := gameManager.CurrBuffer[windowWidth*int(e.pos.y)+int(e.pos.x)-1]
			if charOnNextPosition != e.sprite {
				e.pos.x -= speed
			}
		} else {
		}

		if gameManager.ptrPlayer.pos.y > int(e.pos.y) {

			charOnNextPosition := gameManager.CurrBuffer[windowWidth*int(e.pos.y+1)+int(e.pos.x)]
			if charOnNextPosition != e.sprite {
				e.pos.y += speed
			}
		} else if gameManager.ptrPlayer.pos.y < int(e.pos.y) {

			charOnNextPosition := gameManager.CurrBuffer[windowWidth*int(e.pos.y-1)+int(e.pos.x)]
			if charOnNextPosition != e.sprite {
				e.pos.y -= speed
			}
		} else {
		}
	} else {
		//gameManager.writeToConsole("PLAYER", "NIL")
	}
}

func (enemy *Enemy) draw() {
	characterOnCurrentPosition := gameManager.CurrBuffer[windowWidth*int(enemy.pos.y)+int(enemy.pos.x)]
	if characterOnCurrentPosition != '?' && characterOnCurrentPosition != '&' && characterOnCurrentPosition != ' ' {
		gameManager.killEnemy(enemy.id, enemy.creationId)
	}
	if int(enemy.pos.x) <= windowWidth && int(enemy.pos.y) <= windowHeight {
		gameManager.CurrBuffer[windowWidth*int(enemy.pos.y)+int(enemy.pos.x)] = enemy.sprite
	}
}
