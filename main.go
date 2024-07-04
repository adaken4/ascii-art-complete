package main

import (
	"fmt"
	"os"

	"asciiart/ascii"
	"asciiart/flags"
	"asciiart/justify"
)

func main() {

	options := flags.ParseOptions()

	runeAsciiArtMap, err := ascii.RuneAsciiArtMapCreator("./banners/" + options.Banner + ".txt")
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	artText := ascii.ArtStringBuilder(options.Input, runeAsciiArtMap)
	artText = justify.ArtAligner(options.Align, artText)

	fmt.Print(artText)
}
