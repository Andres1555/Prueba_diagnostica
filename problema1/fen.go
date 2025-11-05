package main

import (
	"fmt"
	"regexp"
)

func isValidFEN(fen string) bool {
	pattern := `^([rnbqkpRNBQKP1-8]{1,8}/){7}[rnbqkpRNBQKP1-8]{1,8} [wb] (-|[KQkq]{1,4}) (-|[a-h][36]) \d+ \d+$`
	match, _ := regexp.MatchString(pattern, fen)
	return match
}

func main() {
	test := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fmt.Println("¿FEN válida?", isValidFEN(test))
}
