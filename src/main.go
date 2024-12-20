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
//  Not a complete vector implementation
// Just here cause pygame and unity
type Vector2 struct {
	x, y int
}


// Level

const windowHeight int = 45

type Level struct {
	sprite []string
}

func (l *Level) draw() {
	fmt.Printf("\033[s")
	for idx, val := range l.sprite {
		fmt.Printf("\033[%d;%dH%s", idx, 0, val)
	}
	fmt.Printf("\033[u")}

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
	//	fmt.Print("\033[2J\033[H")
	for i := 0; i < len(elems); i++ {
		elems[i].draw()
	}
}


// Utility Funcs

func getCursorCords() string {
	ansi_string := "\033[6n"
	return ansi_string
}

func abs(x int) int {
	if x < 0 {
		return 0 - x 
	}
	return x 
}

func contentsOfFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return make([]string, 0), err
	}
	defer file.Close()

	buf := make([]byte, 1)
	data := make([]string, windowHeight)
	var currentLine string
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error Reading file: ", err)
				panic(err)
			}
			break
		}

		char := string(buf[0])

		if char == "\n" {
			data = append(data, string(buf[:n]))
			currentLine = ""			
		} else {
			currentLine += char
		}
	}
	return data, nil
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
		drawable := []Drawable{}

		fmt.Print("\033[2J\033[H")
		fmt.Println("Use arrow keys to move '@'. Press 'q' to quit.")
		fmt.Print("\033[3;3H@")

		buf := make([]byte, 1)


		player := &Player{pos: Vector2{0, 0}, sprite: "&"}
		level := &Level{sprite: make([]string, windowHeight)}
		
		level.sprite, err = contentsOfFile("level.txt")
		if err != nil {
			panic(err)
		}
		level.draw()
		drawable = append(drawable, player)
		//drawable = append(drawable, level)
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
	case "level_editor":	// I dont need this
		fmt.Printf("\033[2J\033[H")

		var level []string = make([]string, 45)
		for idx := range level {			
			level[idx] = strings.Repeat("# ", 95) // 95x47 is how much it takes for the whole screen on my laptop
		}	

		drawSquare := func (lev *[]string, replacementChar string, trueSquare bool, radius int){
			radius = abs(radius)
			centerx, centery := 80, 22

			for idx, val := range level {
				if centery - radius <= idx  && centery + radius >= idx {
					left, right := centerx - radius, centerx + radius
					if trueSquare {
						(*lev)[idx] = val[0:left] + strings.Repeat(replacementChar, radius) + val[right:]
					} else {
						(*lev)[idx] = val[0:left] + strings.Repeat(replacementChar, 2 * radius) + val[right:]
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
				fmt.Println("Error writing to file, " , err)
			}
			
		}

		drawSquare(&level, "  ", false, 20)

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