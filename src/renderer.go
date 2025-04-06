package main

import "regexp"

type Renderer struct{}

type indices struct {
	start int
	end   int
}

func (renderer *Renderer) parseANSI(buf []rune) []indices {
	ANSIRegex := regexp.MustCompile(`\033\[[0-9;]*[a-zA-Z]`)

	var chunks []indices
	matches := ANSIRegex.FindAllStringIndex(string(buf), -1)

	// Convert matches to chunks
	for _, val := range matches {
		chunks = append(chunks, indices{start: val[0], end: val[1]})
	}

	return chunks
}
