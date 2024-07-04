package main

import (
	"fmt"
	"os"

	"asciiart/ascii"
	"asciiart/flags"
)

func main() {

	options := flags.ParseOptions()

	runeAsciiArtMap, err := ascii.RuneAsciiArtMapCreator("./banners/" + options.Banner + ".txt")
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	artText := ascii.ArtStringBuilder(options.Input, runeAsciiArtMap)

	// Print the ASCII art
	fmt.Print(artText)
}
