package main

import (
	"fmt"
	"os"
	"strings"

	"asciiart/ascii"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("provide ascii text as one argument")
		return
	}
	artText, err := ArtBuilder(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(artText)
}

func ArtBuilder(s string) (string, error) {
	asciiMap, err := ascii.RuneAsciiArtMapCreator("thinkertoy.txt")
	if err != nil {
		return "", err
	}
	var result strings.Builder
	for i := 0; i < 8; i++ {
		for _, v := range s {
			artLines := strings.Split(asciiMap[v], "\n")
			result.WriteString(artLines[i])
		}
		result.WriteString("\n")
	}
	return result.String(), nil
}
