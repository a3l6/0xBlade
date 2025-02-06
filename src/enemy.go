package main

type Enemy struct {
	pos        fVector2
	lastPos    fVector2
	player     *Player
	sprite     byte
	vel        Vector2
	damage     uint
	health     int
	id         int
	creationId int
}

func (e *Enemy) collision(obj int) {
	//ptr_obj := gameManager.drawable[obj]

}

// Simple movement towards player
func (e *Enemy) Step() {
	/*if gameManager.ptrPlayer != nil {
		if gameManager.ptrPlayer.pos.y > e.pos.y {
			e.vel.y++
		} else if gameManager.ptrPlayer.pos.y < e.pos.y {
			e.vel.x--
		}

		if gameManager.ptrPlayer.pos.x+1 > e.pos.x {
			e.vel.x++
		} else if gameManager.ptrPlayer.pos.x+1 < e.pos.x {
			e.vel.x--
		}

		e.pos = addVector2(e.pos, e.vel)

		if gameManager.ptrPlayer.pos == e.pos {
			e.pos = addVector2(e.pos, Vector2{x: 1, y: 1})
		}

	}*/
	if gameManager.ptrPlayer != nil {
		const speed = .5
		e.lastPos = e.pos
		if gameManager.ptrPlayer.pos.x > int(e.pos.x) {
			e.pos.x += speed
		} else if gameManager.ptrPlayer.pos.x < int(e.pos.x) {
			e.pos.x -= speed
		} else {
		}

		if gameManager.ptrPlayer.pos.y > int(e.pos.y) {
			e.pos.y += speed
		} else if gameManager.ptrPlayer.pos.y < int(e.pos.y) {
			e.pos.y -= speed
		} else {
		}
	} else {
		//gameManager.writeToConsole("PLAYER", "NIL")
	}

	/*
		coord := gameManager.getCoordinate(Vector2{x: int(e.pos.x), y: int(e.pos.y)})
		if coord == GRENADE {
			gameManager.killEnemy(e.id, e.creationId)
		}*/

}

func (enemy *Enemy) draw() {
	characterOnCurrentPosition := gameManager.CurrBuffer[windowWidth*int(enemy.pos.y)+int(enemy.pos.x)]
	if characterOnCurrentPosition != '?' && characterOnCurrentPosition != '&' && characterOnCurrentPosition != ' ' {
		gameManager.killEnemy(enemy.id, enemy.creationId)
	}
	if int(enemy.pos.x) <= windowWidth && int(enemy.pos.y) <= windowHeight {
		gameManager.CurrBuffer[windowWidth*int(enemy.pos.y)+int(enemy.pos.x)] = enemy.sprite
	}
	/*fmt.Printf("\033[s")
	if enemy.lastPos != (fVector2{x: 0, y: 0}) {
		fmt.Printf("\033[%d;%dH ", int(enemy.lastPos.y), int(enemy.lastPos.x))
	}

	fmt.Printf("\033[%d;%dH%s", int(enemy.pos.y), int(enemy.pos.x), enemy.sprite)
	fmt.Printf("\033[u")*/
}
