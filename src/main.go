package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"
)

const windowHeight uint8 = 45
const windowWidth uint8 = 95 * 2

var gameManager GameManager = GameManager{objects: make([]*Objects, 1)}
var level Level = Level{
	upperBound: 2,
	lowerBound: 42,
	rightBound: 140,
	leftBound:  60,
	sprite:     make([]string, windowHeight),
}
var keymap Keymap = Keymap{up: 'w', down: 's', right: 'd', left: 'a'}

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

	buf := make([]byte, 3)

	handleInputs := func() {
		const fps = 120
		frameDuration := time.Second / fps
		for {
			start := time.Now()

			_, err := os.Stdin.Read(buf)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			gameManager.drawScreen()
			elapsed := time.Since(start)
			sleepTime := frameDuration - elapsed
			if sleepTime > 0 {
				time.Sleep(sleepTime)
			}
		}
	}

	player := Player{pos: Vector2{x: 10, y: 10}, sprite: '&'}
	player.id = gameManager.registerObject(player)

	go handleInputs()
	const fps = 5
	frameDuration := time.Second / fps
	for {
		start := time.Now()

		if buf[0] == 'q' {
			return
		}
		elapsed := time.Since(start)
		sleepTime := frameDuration - elapsed
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}
