package main

import (
	"fmt"
	/*
		"log"
		"os"
		"os/signal"
		"syscall"

		"golang.org/x/term"*/)

type foo struct {
	bar  int
	bazz uint8
}

func main() {

	var ints [100]foo
	fmt.Print(ints)
	/*oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
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
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		//var test [3][4]int
		//fmt.Print(test)
		//fmt.Printf("%T", buf)
		fmt.Print(fmt.Sprintf("\033[6n"))
		fmt.Printf("\033[6n")
	}
	*/
}
