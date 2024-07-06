package main

import (
	"fmt"
	"os"

	"asciiart/ascii"
	"asciiart/flags"
	"asciiart/justify"
	"asciiart/output"
)

func main() {

	// Get terminal flag options and arguments
	options, err := flags.ParseOptions()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	// Set the default banner file
	if options.Banner == "" {
		options.Banner = "rounded"
	}

	// Create a map of each printable Ascii rune to art
	runeAsciiArtMap, err := ascii.RuneAsciiArtMapCreator("./banners/" + options.Banner + ".txt")
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	// Define parameters for Ascii art retrieval
	params := ascii.ArtParams{
		InputText:   options.Input,
		SubString:   options.Substring,
		Colour:      options.Color,
		AsciiArtMap: runeAsciiArtMap,
	}

	// Build art representation of the input text
	artText, err := ascii.ArtStringBuilder(params)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	// Align art representation if specified
	if options.Align != "" {
		artText, err = justify.ArtAligner(options.Align, artText)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
		}
	}

	// Output art representation to a file if provided
	if options.Output != "" {
		err = output.OutputWriter(options.Output, artText)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
		}
	} else {
		// Print art representation to terminal
		fmt.Print(artText)
	}
}
