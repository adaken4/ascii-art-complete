package main

import (
	"fmt"
	"os"

	"asciiart/ascii"
)

func main() {
	// Get command-line arguments
	args := os.Args[1:]

	// Check if the correct number of arguments is provided
	if len(args) != 2 {
		fmt.Println("provide input text and banner name")
		return
	}

	runeAsciiArtMap, err := ascii.RuneAsciiArtMapCreator("./banners/" + args[1] + ".txt")
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	artText := ascii.ArtStringBuilder(args[0], runeAsciiArtMap)

	// Print the ASCII art
	fmt.Print(artText)
}
