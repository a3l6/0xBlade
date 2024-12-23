package main

import (
	"fmt"
	"os"
)

func Abs(x int) int {
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
	var data []string
	var currentLine string
	for {
		_, err := file.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error Reading file: ", err)
				panic(err)
			}
			break
		}

		char := string(buf[0])

		if char == "\n" {
			data = append(data, currentLine)
			currentLine = ""			
		} else {
			currentLine += char
		}
	}
	return data, nil
}
