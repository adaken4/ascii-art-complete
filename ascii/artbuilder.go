package ascii

import (
	"regexp"
	"strings"
)

// ArtStringBuilder generates ASCII art for a given string using a specified banner file.
func ArtStringBuilder(s string, asciiArtMap map[rune]string) string {

	// Create a strings.Builder to build the ASCII art string
	var result strings.Builder
	newlines := regexp.MustCompile(`\\n`)
	s = newlines.ReplaceAllString(s, "\n")

	// Handle newlines accordingly, if input text contains only newlines
	if onlyNewLines(s) {
		result.WriteString(s)
		return result.String()
	}

	inputSlices := strings.Split(s, "\n")

	for _, v := range inputSlices {
		if v == "" {
			result.WriteString("\n")
		} else {
			artString := StringBuilder(v, asciiArtMap)
			result.WriteString(artString)
		}
	}

	// Return the generated ASCII art string
	return result.String()
}

// onlyNewLines checks if a string contains only newline runes
func onlyNewLines(s string) bool {
	for _, v := range s {
		if v != '\n' {
			return false
		}
	}
	return true
}

func StringBuilder(s string, asciiArtMap map[rune]string) string {

	var result strings.Builder

	// Iterate through each line of the ASCII art (8 lines per character)
	for i := 0; i < 8; i++ {
		// Iterate through each character in the input string
		for _, v := range s {
			// Get the ASCII art lines for the current character
			artLines := strings.Split(asciiArtMap[v], "\n")

			// Write the current line of ASCII art for the character
			result.WriteString(artLines[i])
		}

		// Add a newline after each line of ASCII art
		result.WriteString("\n")
	}

	return result.String()
}
