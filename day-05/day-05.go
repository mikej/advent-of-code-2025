package main

import "strings"

func SplitInput(input string) ([]string, []string) {
	parts := strings.SplitN(input, "\n\n", 2)
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}
