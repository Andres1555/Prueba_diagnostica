package main

import (
	"fmt"
	"regexp"
)

func main() {
	tests := map[string]string{
		"IP":         `^(?:\d{1,3}\.){3}\d{1,3}$`,
		"Email":      `^[\w\.-]+@[\w\.-]+\.\w{2,}$`,
		"Científica": `^[+-]?\d+(\.\d+)?[eE][+-]?\d+$`,
	}

	inputs := []string{"192.168.0.1", "user@mail.com", "-3.2e+10"}

	for _, input := range inputs {
		for label, pattern := range tests {
			match, _ := regexp.MatchString(pattern, input)
			if match {
				fmt.Printf("'%s' es %s\n", input, label)
			}
		}
	}
}
