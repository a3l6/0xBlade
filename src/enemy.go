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

	gameManager.writeToConsole("POS", fmt.Sprintf("%d %d", gameManager.ptrPlayer.pos.x, gameManager.ptrPlayer.pos.y))
	//gameManager.console["POS"] = fmt.Sprintf("%d %d", gameManager.ptrPlayer.pos.x, gameManager.ptrPlayer.pos.y)
	//gameManager.console["2POS"] = fmt.Sprintf("%d %d", e.pos.x, e.pos.y)
	coord := gameManager.getCoordinate(e.pos)
	if coord == GRENADE {
		gameManager.writeToConsole("DIED", "TRUE")
		//gameManager.console["DIED"] = "TRUE"
	}

}

func (enemy Enemy) draw() {
	fmt.Printf("\033[s")
	fmt.Printf("\033[%d;%dH%s", enemy.pos.y, enemy.pos.x, enemy.sprite)
	fmt.Printf("\033[u")

}
