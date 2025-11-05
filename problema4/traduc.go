package main

import (
	"fmt"
	"strings"
)

var cKeywords = map[string]string{
	"int":    "entero",
	"float":  "flotante",
	"char":   "carácter",
	"return": "retornar",
	"if":     "si",
	"else":   "sino",
	"for":    "para",
	"while":  "mientras",
	"struct": "estructura",
	"void":   "vacío",
}

func translateCCode(code string) string {
	words := strings.Fields(code)
	for i, word := range words {
		if val, ok := cKeywords[word]; ok {
			words[i] = val
		}
	}
	return strings.Join(words, " ")
}

func main() {
	cCode := "int main() { return 0; }"
	fmt.Println("Traducido:", translateCCode(cCode))
}
