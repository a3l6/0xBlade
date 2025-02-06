package main

type Grenade struct {
	pos         Vector2
	vel         Vector2
	sprite      string
	trailSprite string
	step        int // between 1-4
	amplitude   int
	id          int
	creationID  uint8
}

func (g *Grenade) Step() {}

// TODO: Mke grenade explosion
func (g *Grenade) draw() {
	//fmt.Printf("\033[s")
	sprite := g.trailSprite

	// TODO: Change to stepable code
	// LEGACY CODE
	fps := 10 // FPS here to make it run slower
	switch {
	case g.step == 1*fps:
		g.pos.y -= 1 * g.amplitude
		g.pos.x++
		g.step++
		/*fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x) // All of these remove the old character
		g.pos.y -= 1 * g.amplitude
		g.pos.x++
		g.step++*/
	case g.step == 2*fps:
		//fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		g.pos.y -= 2 * g.amplitude
		g.pos.x++
		g.step++
	case g.step == 3*fps:
		//fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		g.pos.y += 2 * g.amplitude
		g.pos.x++
		g.step++
	case g.step == 4*fps:
		//fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		g.pos.y += 1 * g.amplitude
		g.pos.x++
		sprite = g.sprite
		g.step++
	case g.step >= 5*fps && g.step < 7*fps:
		// TODO: Make grenade explosion random
		// Just draws the thing
		// Empty: "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "
		//new_sprite := "\0331A   !#\033[1B\033[4D$#@#$\033[1B\033[5D #@$%$#$#\033[1B\033[5D!@#"
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x+3:], []byte{'!', '#'})
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x+2:], []byte{'$', '#', '@', '#', '$'})
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x+1:], []byte{'#', '@', '$', '%', '$', '#', '$', '#'})
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x+3:], []byte{'!', '@', '#'})
		// Set coordinate
		/*gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: 1}), GRENADE)

		gameManager.setCoordinate(g.pos, GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: 0}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: 0}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 0}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: 0}), GRENADE)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 5, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 6, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 7, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vecto22{x: 8, y: -1}), GRENADE)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: -2}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: -2}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 5, y: -2}), GRENADE)
		*/
		//sprite = new_sprite
		g.step++
	case g.step >= 7*fps && g.step < 8*fps:

		//new_sprite := "\0331A   ! \033[1B\033[4D @#@ \033[1B\033[5D  @#$%   \033[1B\033[5D * "
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x+2:], []byte{'!'})
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x+4:], []byte{'@', '#', '@'})
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x+3:], []byte{'@', '#', '$', '%'})
		copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x+3:], []byte{'*'})

		// Reset
		/*gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: 1}), NOT_FOUND)

		gameManager.setCoordinate(g.pos, NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: 0}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: 0}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 0}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: 0}), NOT_FOUND)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 5, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 6, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 7, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 8, y: -1}), NOT_FOUND)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: -2}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: -2}), NOT_FOUND)

		// Set coordinate
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: 0}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: 0}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 0}), GRENADE)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: -1}), GRENADE)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: -1}), GRENADE)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 5, y: -2}), GRENADE)
		*/
		//sprite = new_sprite`
		g.step++
	case g.step == 8*fps:
		/*new_sprite := "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: 0}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: 0}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: 0}), NOT_FOUND)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 1, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 2, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 3, y: -1}), NOT_FOUND)
		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 4, y: -1}), NOT_FOUND)

		gameManager.setCoordinate(addVector2(g.pos, Vector2{x: 5, y: -2}), NOT_FOUND)

		sprite = new_sprite */
		g.step++
	case g.step == 12*fps:
		//gameManager.setCoordinate(g.pos, NOT_FOUND)
		//fmt.Printf("\033[%d;%dH ", g.pos.y, g.pos.x)
		gameManager.deleteObject(g.id, g.creationID) // kill self
	default:
		g.step++
	}
	//gameManager.setCoordinate(g.pos, GRENADE)
	copy(gameManager.CurrBuffer[windowWidth*g.pos.y+g.pos.x:], []byte(sprite))
	//fmt.Printf("\033[%d;%dH%s", g.pos.y, g.pos.x, sprite)
	//fmt.Printf("\033[u")
}
