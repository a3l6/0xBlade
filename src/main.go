package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

const windowHeight int = 45
const windowWidth int = 95 * 2

var gameManager GameManager = GameManager{
	drawable:            make(map[int]*GameObject),
	console:             make(map[string]string),
	occupiedCoordinates: make(map[Vector2]int),
}
var keymap Keymap = Keymap{up: 'w', down: 's', left: 'a', right: 'd', aimUp: 'i', aimDown: 'k', aimLeft: 'j', aimRight: 'l'}
var level *Level = &Level{
	upperBound: 2,
	lowerBound: 42,
	rightBound: 140,
	leftBound:  60,
}

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

	mode := "main" // main or level_editor

	switch mode {
	case "main":

		buf := make([]byte, 1)

		player := &Player{pos: Vector2{61, 2}, sprite: "&", keymap: keymap}

		sprite, err := contentsOfFile("src/level.txt")
		if err != nil {
			panic(err)
		}

		level.sprite = strings.Join(sprite, "\r\n")

		/*fileName := "level1.txt"
		data := strings.Join(level.sprite, "\n")
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Cannot create file! ", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println("Error writing to file, ", err)
		}*/

		handleInputs := func() {
			frameDuration := time.Second / 120
			for {
				start := time.Now()
				_, err := os.Stdin.Read(buf)
				if err != nil {
					term.Restore(int(os.Stdin.Fd()), oldState)

					log.Fatal(err)
					panic(err)
				}
				player.move(buf)

				//gameManager.console["buf"] = "\"" + string(buf) + "\""
				gameManager.writeToConsole("buf", fmt.Sprintf("\"%s\"", string(buf)))
				if buf[0] == ' ' {
					err = gameManager.createNewGrenade(player.pos)
					if err != nil {
					}
				}

				elapsed := time.Since(start)
				sleepTime := frameDuration - elapsed
				if sleepTime > 0 {
					time.Sleep(sleepTime)
				}
			}

		}

		// Needs to run faster to update all the time
		go func() {
			frameDuration := time.Second / 120
			for {
				start := time.Now()
				gameManager.drawScreen()

				elapsed := time.Since(start)
				sleepTime := frameDuration - elapsed
				if sleepTime > 0 {
					time.Sleep(sleepTime)
				}
			}
		}()

		level.draw()

		//level.id = gameManager.registerAsObject(level)
		player.id = gameManager.registerAsObject(player)

		gameManager.createNewEnemy(Vector2{70, 10}, player)

		const fps = 10
		frameDuration := time.Second / fps

		go handleInputs()

		for {
			start := time.Now()
			gameManager.ptrPlayer = player
			gameManager.StepAll()

			if buf[0] == 'q' {
				return
			}

			if buf[0] == 'r' {
				writeToFile("output1.txt", fmt.Sprint(gameManager.occupiedCoordinates))
			}

			elapsed := time.Since(start)
			sleepTime := frameDuration - elapsed
			if sleepTime > 0 {
				time.Sleep(sleepTime)
			}
		}

	case "level_editor": // I dont need this
		fmt.Printf("\033[2J\033[H")

		var level []string = make([]string, 45)
		for idx := range level {
			level[idx] = strings.Repeat("# ", 95) // 95x47 is how much it takes for the whole screen on my laptop
		}

		drawSquare := func(lev *[]string, replacementChar string, trueSquare bool, radius int) {
			radius = Abs(radius)
			centerx, centery := 80, 22

			for idx, val := range level {
				if centery-radius <= idx && centery+radius >= idx {
					left, right := centerx-radius, centerx+radius
					if trueSquare {
						(*lev)[idx] = val[0:left] + strings.Repeat(replacementChar, radius) + val[right:]
					} else {
						(*lev)[idx] = val[0:left] + strings.Repeat(replacementChar, 2*radius) + val[right:]
					}
				}
			}

			// Save to file

			fileName := "level.txt"
			data := strings.Join(level, "\n")
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Cannot create file! ", err)
				return
			}
			defer file.Close()

			_, err = file.WriteString(data)

			if err != nil {
				fmt.Println("Error writing to file, ", err)
			}

		}

		drawSquare(&level, "  ", false, 20)

		// made it only draw when needed cause computer couldn't keep up when moving the cursor fast
		drawLevel := func() {
			fmt.Printf("\033[s")
			for idx, val := range level {
				fmt.Printf("\033[%d;%dH%s", idx, 0, val)
			}
			fmt.Printf("\033[u")
		}

		drawLevel()

		buf := make([]byte, 3)

		//console := []string{}
		//x, y := 0, 0
		for {

			//fmt.Printf("\033[%d;%dH", y, x)
			_, err = os.Stdin.Read(buf)
			if err != nil {
				log.Fatal(err)
			}

			switch buf[0] {
			case 'q':
				return
			case ' ':
				fmt.Printf("\033[s")
				fmt.Printf("\033[47;0HCONSOLE: %s", "\033[6n")
				fmt.Printf("\033[u")
				drawLevel()
			case '\033':
				if buf[1] == '[' {
					fmt.Printf("\033[1%s", string(buf[2]))
				}
			}
			//console = append(console, fmt.Sprintf("X: %d  Y: %d", x, y))e

		}

	}
}
