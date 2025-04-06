package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/a3l6/libtui"
	"golang.org/x/term"
)

const windowHeight int = 45
const windowWidth int = 95 * 2

var gameManager GameManager = GameManager{
	drawable:       make(map[int]*GameObject),
	console:        make(map[string]string),
	difficultySeed: 120,
}
var keymap Keymap = Keymap{up: 'w', down: 's', left: 'a', right: 'd', aimUp: 'i', aimDown: 'k', aimLeft: 'j', aimRight: 'l'}
var level *Level = &Level{
	upperBound: 2,
	lowerBound: 42,
	rightBound: 140,
	leftBound:  60,
}

var screenBuffer ScreenBuffer
var Writer = bufio.NewWriterSize(&screenBuffer, windowHeight*windowWidth)

var renderer Renderer

var mode string = "menu"

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
	for {
		switch mode {
		case "menu":
			width, height, err := term.GetSize(int(os.Stderr.Fd()))
			if err != nil {
				width := windowWidth
				height := windowHeight
				fmt.Printf("Error: %d %d", width, height)
			}

			splash_screen := strings.Split(
				"▒█████  ▒██   ██▒ ▄▄▄▄    ██▓    ▄▄▄      ▓█████▄ ▓█████\r\n"+
					"▒██▒  ██▒▒▒ █ █ ▒░▓█████▄ ▓██▒   ▒████▄    ▒██▀ ██▌▓█   ▀ \r\n"+
					"▒██░  ██▒░░  █   ░▒██▒ ▄██▒██░   ▒██  ▀█▄  ░██   █▌▒███   \r\n"+
					"▒██   ██░ ░ █ █ ▒ ▒██░█▀  ▒██░   ░██▄▄▄▄██ ░▓█▄   ▌▒▓█  ▄ \r\n"+
					"░ ████▓▒░▒██▒ ▒██▒░▓█  ▀█▓░██████▒▓█   ▓██▒░▒████▓ ░▒████▒\r\n"+
					"░ ▒░▒░▒░ ▒▒ ░ ░▓ ░░▒▓███▀▒░ ▒░▓  ░▒▒   ▓▒█░ ▒▒▓  ▒ ░░ ▒░ ░\r\n"+
					"  ░ ▒ ▒░ ░░   ░▒ ░▒░▒   ░ ░ ░ ▒  ░ ▒   ▒▒ ░ ░ ▒  ▒  ░ ░  ░\r\n"+
					"░ ░ ░ ▒   ░    ░   ░    ░   ░ ░    ░   ▒    ░ ░  ░    ░   \r\n"+
					"    ░ ░   ░    ░   ░          ░  ░     ░  ░   ░       ░  ░\r\n"+
					"                        ░                   ░              \r\n",
				"\r\n")

			settings_modal := libtui.Modal{Width: 10, Height: 20, Position: libtui.Vector2{X: 20, Y: 20}}

			var buttons [3]libtui.Button
			buttons[0] = libtui.Button{
				Width:  20,
				Height: 1,
				Align:  libtui.AlignLeft,
				Position: libtui.Vector2{
					// 20 is width, 15 is just offset for sep
					X: width/2 - 20 - 15,
					Y: height / 2},
				Value: "Start Game     spc",
				Key:   ' ',
				Callback: func() {
					mode = "main"
				},
			}
			buttons[1] = libtui.Button{
				Width:  20,
				Height: 1,
				Align:  libtui.AlignLeft,
				Position: libtui.Vector2{
					X: width/2 - 10,
					Y: height / 2},
				Value: "Settings         s",
				Key:   's',
				Callback: func() {
					settings_modal.Active = true
				},
			}
			buttons[2] = libtui.Button{
				Width:  20,
				Height: 1,
				Align:  libtui.AlignLeft,
				Position: libtui.Vector2{
					X: width/2 + 15,
					Y: height / 2},
				Value:    "Quit             q",
				Key:      'q',
				Callback: func() {},
			}

			const fps = 10
			const frameDuration = time.Second / fps
			buf := make([]byte, 1)

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
					elapsed := time.Since(start)
					sleepTime := frameDuration - elapsed
					if sleepTime > 0 {
						time.Sleep(sleepTime)
					}
					if mode != "menu" {
						break
					}

				}

			}
			go handleInputs()

			for {
				start := time.Now()

				for idx, row := range splash_screen {
					y := idx + 10 //width * (idx + (height/2 - len(splash_screen)) - 3) // 3 is just because
					x := width/2 - len(splash_screen[0])/4
					fmt.Fprintf(Writer, "\033[%d;%dH%s", y, x, row)
				}

				for _, btn := range buttons {
					if buf[0] == byte(btn.Key) {
						btn.Callback()
					}

					rendered, err := btn.RenderToArrRunes()
					if err != nil {
						log.Fatal(err)
					}

					fmt.Fprintf(Writer, "\033[%d;%dH%c", btn.Position.Y, btn.Position.X, rendered)
				}

				if buf[0] == 'q' {
					return
				}

				if settings_modal.Active {
					rendered, err := settings_modal.RenderToArrRunes()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Fprintf(Writer, "\033[%d;%dH%c", settings_modal.Position.Y, settings_modal.Position.X, rendered)
				}

				elapsed := time.Since(start)
				sleepTime := frameDuration - elapsed
				if sleepTime > 0 {
					time.Sleep(sleepTime)
				}

				if mode != "menu" {
					break
				}

				Writer.Flush()
				fmt.Fprintf(Writer, "\033[2J")
			}

		case "main":

			buf := make([]byte, 1)

			player := &Player{pos: Vector2{61, 2}, sprite: "&", keymap: keymap}

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
						err = gameManager.createNewGrenade(player.pos, player.last_direction)
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

			//level.id = gameManager.registerAsObject(level)
			player.id = gameManager.registerAsObject(player)

			gameManager.createNewEnemy(Vector2{70, 10})

			const fps = 10
			const frameDuration = time.Second / fps

			go handleInputs()

			for {
				start := time.Now()
				// TODO: convert to once??
				gameManager.ptrPlayer = player

				gameManager.StepAll()

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
}
