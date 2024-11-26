package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func main() {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

	fmt.Print("\033[2J\033[H")
	fmt.Println("Use arrow keys to move '@'. Press 'q' to quit.")
	fmt.Print("\033[3;3H@")

	posX, posY := 3, 3
	buf := make([]byte, 3)

	for {
		os.Stdin.Read(buf)
		switch buf[0] {
		case 'q':
			fmt.Print("\033[2J\033[H")
			fmt.Println("Goodbye!")
			return
		case '\033': // Escape sequence
			if buf[1] == '[' {
				switch buf[2] {
				case 'A': // Up
					posY--
				case 'B': // Down
					posY++
				case 'C': // Right
					posX++
				case 'D': // Left
					posX--
				}
			}
		}
		fmt.Print("\033[2J\033[H")
		fmt.Printf("\033[%d;%dH@", posY, posX)
	}
}