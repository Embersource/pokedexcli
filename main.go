package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	output := strings.Fields(lowered)
	return output
}

func main() {
	fmt.Println("Hello, World!")
}
