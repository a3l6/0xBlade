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
	direction   uint8
}

func (g *Grenade) Step() {}

// Copies the grenades sprite to the buffer. All panics are ignored.
func (g *Grenade) draw() {
	sprite := g.trailSprite

	/*

			UP


		       !@!#!
		       #@!@$#@#
			@$#$@
			 !#!
			 #

		       !@!
		        @!@#
			 $#@
			 !#
			 #



			DOWN
			 #
		        !#!
		       @$#$@
		      #@#$@!@#
		      !#!@#

			 #
		         #!
		        $#$@
		       #!@#
		      !#!


			RIGHT
			opp of left

	*/

	// TODO: Change to stepable code
	// LEGACY CODE
	fps := 10 // FPS here to make it run slower

	defer func() {
		if err := recover(); err != nil {
			// When copying the grenade sprite to the buffer it will panic if out of range
			// This is fine, just get rid of the grenade because it's usually just off the screen
			gameManager.deleteObject(g.id, g.creationID)
		}
	}()

	switch g.direction {
	case DIRECTION_UP:
		switch {
		case g.step == 1*fps:
			g.pos.y -= 1 * g.amplitude
			g.pos.x++
			g.step++
		case g.step == 2*fps:
			g.pos.y -= 2 * g.amplitude
			g.step++
		case g.step == 3*fps:
			g.pos.y += 2 * g.amplitude
			g.step++
		case g.step == 3*fps:
			g.step++
		case g.step >= 5*fps && g.step < 7*fps:
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-4)+g.pos.x-2:], []rune{'!', '@', '!', '#', '!'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-3)+g.pos.x-2:], []rune{'#', '@', '!', '@', '$', '#', '@', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-2)+g.pos.x-1:], []rune{'@', '$', '#', '$', '@'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-1)+g.pos.x:], []rune{'!', '#', '!'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x:], []rune{'#'})
			g.step++
		case g.step >= 7*fps && g.step < 8*fps:
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-3)+g.pos.x-2:], []rune{'!', '@', '!'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-3)+g.pos.x-1:], []rune{'@', '!', '@', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-2)+g.pos.x:], []rune{'$', '#', '@'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y-1)+g.pos.x:], []rune{'!', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x:], []rune{'#'})
			g.step++
		case g.step == 8*fps:
			g.step++
		case g.step == 12*fps:
			gameManager.deleteObject(g.id, g.creationID) // kill self
		default:
			g.step++
		}

	case DIRECTION_DOWN:

		switch {
		case g.step == 1*fps:
			g.pos.y += 1 * g.amplitude
			g.pos.x++
			g.step++
		case g.step == 2*fps:
			g.pos.y += 2 * g.amplitude
			g.step++
		case g.step == 3*fps:
			g.pos.y -= 2 * g.amplitude
			g.step++
		case g.step == 3*fps:
			g.step++
		case g.step >= 5*fps && g.step < 7*fps:
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x:], []rune{'#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x:], []rune{'!', '#', '!'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x-1:], []rune{'@', '$', '#', '$', '@'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x-2:], []rune{'#', '@', '!', '@', '$', '#', '@', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+4)+g.pos.x-2:], []rune{'!', '@', '!', '#', '!'})
			g.step++
		case g.step >= 7*fps && g.step < 8*fps:

			/*

				       !@!
				        @!@#
					 $#@
					 !#
					 #

			*/
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x:], []rune{'#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x:], []rune{'!', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x:], []rune{'$', '#', '@'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x-1:], []rune{'@', '!', '@', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x-2:], []rune{'!', '@', '!'})
			g.step++
		case g.step == 8*fps:
			g.step++
		case g.step == 12*fps:
			gameManager.deleteObject(g.id, g.creationID) // kill self
		default:
			g.step++
		}
	case DIRECTION_RIGHT:
		switch {
		case g.step == 1*fps:
			g.pos.y -= 1 * g.amplitude
			g.pos.x++
			g.step++
		case g.step == 2*fps:
			g.pos.y -= 2 * g.amplitude
			g.pos.x++
			g.step++
		case g.step == 3*fps:
			g.pos.y += 2 * g.amplitude
			g.pos.x++
			g.step++
		case g.step == 4*fps:
			g.pos.y += 1 * g.amplitude
			g.pos.x++
			sprite = g.sprite
			g.step++
		case g.step >= 5*fps && g.step < 7*fps:
			// TODO: Make grenade explosion random
			// Empty: "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "
			//new_sprite := "\0331A   !#\033[1B\033[4D$#@#$\033[1B\033[5D #@$%$#$#\033[1B\033[5D!@#"
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x+3:], []rune{'!', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x+2:], []rune{'$', '#', '@', '#', '$'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x+1:], []rune{'#', '@', '$', '%', '$', '#', '$', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x+3:], []rune{'!', '@', '#'})
			g.step++
		case g.step >= 7*fps && g.step < 8*fps:

			//new_sprite := "\0331A   ! \033[1B\033[4D @#@ \033[1B\033[5D  @#$%   \033[1B\033[5D * "
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x+2:], []rune{'!'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x+4:], []rune{'@', '#', '@'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x+3:], []rune{'@', '#', '$', '%'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x+3:], []rune{'*'})
			g.step++
		case g.step == 8*fps:
			g.step++
		case g.step == 12*fps:
			gameManager.deleteObject(g.id, g.creationID) // kill self
		default:
			g.step++
		}

	case DIRECTION_LEFT:
		switch {
		case g.step == 1*fps:
			g.pos.y -= 1 * g.amplitude
			g.pos.x--
			g.step++
		case g.step == 2*fps:
			g.pos.y -= 2 * g.amplitude
			g.pos.x--
			g.step++
		case g.step == 3*fps:
			g.pos.y += 2 * g.amplitude
			g.pos.x--
			g.step++
		case g.step == 4*fps:
			g.pos.y += 1 * g.amplitude
			g.pos.x--
			sprite = g.sprite
			g.step++
		case g.step >= 5*fps && g.step < 7*fps:
			// TODO: Make grenade explosion random
			// Empty: "\0331A     \033[1B\033[4D     \033[1B\033[5D         \033[1B\033[5D   "
			//new_sprite := "\0331A   !#\033[1B\033[4D$#@#$\033[1B\033[5D #@$%$#$#\033[1B\033[5D!@#"
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x-5:], []rune{'!', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x-7:], []rune{'$', '#', '@', '#', '$'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x-9:], []rune{'#', '@', '$', '%', '$', '#', '$', '#'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x-6:], []rune{'!', '@', '#'})
			g.step++
		case g.step >= 7*fps && g.step < 8*fps:

			//new_sprite := "\0331A   ! \033[1B\033[4D @#@ \033[1B\033[5D  @#$%   \033[1B\033[5D * "
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y)+g.pos.x-3:], []rune{'!'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+1)+g.pos.x-7:], []rune{'@', '#', '@'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+2)+g.pos.x-7:], []rune{'@', '#', '$', '%'})
			copy(gameManager.CurrBuffer[windowWidth*(g.pos.y+3)+g.pos.x-4:], []rune{'*'})
			g.step++
		case g.step == 8*fps:
			g.step++
		case g.step == 12*fps:
			gameManager.deleteObject(g.id, g.creationID) // kill self
		default:
			g.step++
		}
	}
	copy(gameManager.CurrBuffer[windowWidth*g.pos.y+g.pos.x:], []rune(sprite))
}
