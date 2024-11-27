package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

// Vector
// Probably not a complete vector implementation
type Vector2 struct {
	x, y int
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
		// Universe not as expected
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


	drawable := []Drawable{}

	fmt.Print("\033[2J\033[H")
	fmt.Println("Use arrow keys to move '@'. Press 'q' to quit.")
	fmt.Print("\033[3;3H@")

	buf := make([]byte, 1)


	player := &Player{pos: Vector2{0, 0}, sprite: "|--O--|"}
	drawable = append(drawable, player)
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
}