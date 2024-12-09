package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/term"
)

// Vector
// Probably not a complete vector implementation
type Vector2 struct {
	x, y int
}


// Level
type Level struct {
	sprite string
}

func (l *Level) loadLevel(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not load level file!")
		return
	}
	l.sprite = string(data)
}

func (l *Level) draw() {
	fmt.Printf("\033[0;0H%s", l.sprite)
}

// Player
type Player struct {
	pos Vector2
	sprite string
}

func (p *Player) move(char uint8) {
	switch char {
	case 'w':
		p.pos.y--
	case 's':
		p.pos.y++
	case 'a':
		p.pos.x--
	case 'd':
		p.pos.x++
	}

}

func (p *Player) draw() {
	fmt.Printf("\033[%d;%dH%s", p.pos.y, p.pos.x, p.sprite)
}


// Drawing 
type Drawable interface {
	draw()
}

func drawAll(elems []Drawable) {
	fmt.Print("\033[2J\033[H")
	for i := 0; i < len(elems); i++ {
		elems[i].draw()
	}
}


// Misc 

func getCursorCords() string {
	ansi_string := "\033[6n"
	return ansi_string
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


	
	mode := "level_editor" // main or level_editor


	switch mode {
	case "main":
		drawable := []Drawable{}

		fmt.Print("\033[2J\033[H")
		fmt.Println("Use arrow keys to move '@'. Press 'q' to quit.")
		fmt.Print("\033[3;3H@")

		buf := make([]byte, 1)


		player := &Player{pos: Vector2{0, 0}, sprite: "&"}
		level := &Level{sprite: ""}
		l := `###############################################################
		###############################################################`
		level.sprite = l
		drawable = append(drawable, player)
		drawable = append(drawable, level)
		for {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			player.move(buf[0])
			drawAll(drawable)
			
			if buf[0] == 'q' {
				return
			}
		}
	case "level_editor":
		fmt.Printf("\033[2J\033[H")

		var level [45]string
		for idx := range level {			
			level[idx] = strings.Repeat("# ", 95) // 95x47 is how much it takes for the whole screen on my laptop
		}	
		
		// made it only draw when needed cause computer couldn't keep up when moving the cursor fast
		drawLevel := func () {
			fmt.Printf("\033[s")
			for idx, val := range level {
				fmt.Printf("\033[%d;%dH%s", idx, 0, val)
			}
			fmt.Printf("\033[u")
		}

		drawLevel()

		buf := make([]byte, 3)
		
		
		console := []string{}
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
				console = append(console, getCursorCords())
				drawLevel()
			case '\033':
				if buf[1] == '[' {
					fmt.Printf("\033[1%s", string(buf[2]))

				}
			}
			//console = append(console, fmt.Sprintf("X: %d  Y: %d", x, y))
			if len(console) > 0 {
				fmt.Printf("\033[s")
				fmt.Printf("\033[47;0H%s", console[0])
				fmt.Printf("\033[u")
			}
			
		}


	}
}