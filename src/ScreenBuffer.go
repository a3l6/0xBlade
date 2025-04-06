package main

import (
	"fmt"
	"unicode/utf8"
)

type ScreenBuffer struct {
	buf     []rune
	prevbuf []rune
}

func (b *ScreenBuffer) Write(curr []byte) (n int, err error) {
	i := 0
	for i < len(curr) {
		// Detect ANSI escape sequence
		if curr[i] == 0x1B && i+1 < len(curr) && curr[i+1] == '[' {
			j := i + 2
			// Allow numbers and semicolons in escape sequence
			for j < len(curr) && ((curr[j] >= '0' && curr[j] <= '9') || curr[j] == ';') {
				j++
			}
			// Include final ANSI command character (like 'H', 'J', 'm', etc.)
			if j < len(curr) {
				j++
			}
			// Print raw ANSI sequence
			fmt.Print(string(curr[i:j]))
			i = j
			continue
		}

		// Otherwise decode UTF-8 rune
		r, size := utf8.DecodeRune(curr[i:])
		if r == utf8.RuneError {
			i += size
			continue
		}

		idx := i
		y := idx / windowWidth
		x := idx % windowWidth

		fmt.Printf("\x1B[%d;%dH%c", y+1, x+1, r)
		i += size
	}

	b.prevbuf = b.buf
	return len(curr), nil
}

/*runes := []rune(string(curr))

for idx, val := range runes {
	y := idx / windowWidth
	x := idx % windowWidth

	if idx == len(curr) {
		break
	}

	fmt.Printf("\x1B[%d;%dH%c", y+1, x+1, val)
	/*if len(b.prevbuf) == 0 || idx > len(b.prevbuf)-1 {
		fmt.Printf("\033[%d;%dH%s", y+1, x+1, string(val))
		continue
	}

	if rune(val) == b.prevbuf[idx] {
		fmt.Printf("\033[%d;%dH%s", y+1, x+1, string(val))
	}*/

//b.prevbuf = b.buf
//return len(curr), nil */
