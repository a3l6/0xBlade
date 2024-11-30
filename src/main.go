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


func main() {
	// Initialize terminal shit
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
		level := strings.Repeat("# ", 95 * 47)
		fmt.Print(level) // 95x47 is how much it takes for the whole screen on my laptop
		
		buf := make([]byte, 3)

		x, y := 0, 0
		for {
			fmt.Printf("\033[2J\033[H")
			fmt.Print(level)
			fmt.Printf("\033[%d;%dH", y, x)
			_, err = os.Stdin.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			switch buf[0] {
			case 'q':
				return
			case '\033':
				if buf[1] == '[' {
					switch buf[2]{
					case 'A':
						y--
					case 'B':
						y++
					case 'C':
						x++
					case 'D':
						x--
					}
				}
			}
			
		}


	}
}