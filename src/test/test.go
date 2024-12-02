package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	buf := make([]byte, 3)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		var test [3][4]int
		fmt.Print(test)
		fmt.Printf("%T", buf)
	}
}