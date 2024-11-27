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
		fmt.Printf("%T", buf)
	}
}