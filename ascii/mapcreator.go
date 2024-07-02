package ascii

import (
	"bufio"
	"os"
	"strings"
)

// RuneAsciiArtMapCreator reads a banner file and creates a map of runes
// to their ASCII art representations.
func RuneAsciiArtMapCreator(bannerPath string) (map[rune]string, error) {
	// Open the banner file
	banner, err := os.Open(bannerPath)
	if err != nil {
		return nil, err
	}
	defer banner.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(banner)
	var art string
	printableRune := ' '
	runeArtMap := make(map[rune]string)

	// Iterate through the file and build the rune-to-ASCII art map
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if art != "" {
				runeArtMap[printableRune] = art
				art = ""
				printableRune++
			}
		} else {
			// Handle Windows-style line endings
			if strings.ContainsRune(line, '\r') {
				line = strings.TrimSuffix(line, "\r") + "\n"
				art += line
			} else {
				art += line + "\n"
			}
		}
	}

	// Add the last ASCII art to the map
	if art != "" {
		runeArtMap[printableRune] = art
	}

	return runeArtMap, nil
}
