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

// Keymap
type Keymap struct {
	up    uint8
	down  uint8
	left  uint8
	right uint8

	aimUp    uint8
	aimDown  uint8
	aimLeft  uint8
	aimRight uint8
}

var gameManager GameManager = GameManager{
	drawable: make(map[int]*Drawable),
	console:  make(map[string]string),
	grenades: make([]Grenade, 0),
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

		//gameManager := GameManager{drawable: make(map[int]*Drawable)}
		//fmt.Print("\033[2J\033[H")
		//fmt.Println("Use arrow keys to move '@'. Press 'q' to quit.")
		//fmt.Print("\033[3;3H@")

		buf := make([]byte, 1)

		level := &Level{sprite: make([]string, windowHeight), upperBound: 2, lowerBound: 42, rightBound: 140, leftBound: 60}
		keymap := Keymap{up: 'w', down: 's', left: 'a', right: 'd', aimUp: 'i', aimDown: 'k', aimLeft: 'j', aimRight: 'l'}
		player := &Player{pos: Vector2{61, 2}, sprite: "&", l: level, keymap: keymap}

		gameManager.createNewGrenade(Vector2{65, 4})

		sprite, err := contentsOfFile("src/level.txt")
		if err != nil {
			panic(err)
		}
		level.sprite = sprite

		fileName := "level1.txt"
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
		}

		handleInputs := func() {
			frameDuration := time.Second / 120
			for {
				start := time.Now()
				_, err := os.Stdin.Read(buf)
				if err != nil {
					log.Fatal(err)
					panic(err)
				}
				player.move(buf[0])
				if buf[0] == ' ' {
					gameManager.createNewGrenade(player.pos)
				}

				elapsed := time.Since(start)
				sleepTime := frameDuration - elapsed
				if sleepTime > 0 {
					time.Sleep(sleepTime)
				}
			}

		}

		level.draw()
		//drawable = append(drawable, level)
		//drawable = append(drawable, player)
		//drawable = append(drawable, &grenade)
		//gameManager.registerAsDrawable(player, &grenade)

		gameManager.registerAsObject(player)
		const fps = 10
		frameDuration := time.Second / fps

		go handleInputs()
		for {
			start := time.Now()

			gameManager.drawScreen()

			if buf[0] == 'q' {
				return
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
