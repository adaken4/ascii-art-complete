package ascii

import "strings"

// ArtStringBuilder generates ASCII art for a given string using a specified banner file.
func ArtStringBuilder(s string, path string) (string, error) {
	// Create a rune-to-ASCII art map using the specified banner file
	asciiMap, err := RuneAsciiArtMapCreator(path)
	if err != nil {
		return "", err
	}

	// Create a strings.Builder to build the ASCII art string
	var result strings.Builder

	// Iterate through each line of the ASCII art (8 lines per character)
	for i := 0; i < 8; i++ {
		// Iterate through each character in the input string
		for _, v := range s {
			// Get the ASCII art lines for the current character
			artLines := strings.Split(asciiMap[v], "\n")

			// Write the current line of ASCII art for the character
			result.WriteString(artLines[i])
		}

		// Add a newline after each line of ASCII art
		result.WriteString("\n")
	}

	// Return the generated ASCII art string
	return result.String(), nil
}
