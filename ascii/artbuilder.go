package ascii

import (
	"regexp"
	"strings"

	"asciiart/color"
)

// ArtStringBuilder generates ASCII art for a given string using a specified banner file.
func ArtStringBuilder(inputText, subString, colour string, asciiArtMap map[rune]string) string {

	// Create a strings.Builder to build the ASCII art string
	var result strings.Builder
	newlines := regexp.MustCompile(`\\n`)
	inputText = newlines.ReplaceAllString(inputText, "\n")

	// Handle newlines accordingly, if input text contains only newlines
	if onlyNewLines(inputText) {
		result.WriteString(inputText)
		return result.String()
	}

	inputSlices := strings.Split(inputText, "\n")

	for _, v := range inputSlices {
		if v == "" {
			result.WriteString("\n")
		} else {
			artString := StringBuilder(v, subString, colour, asciiArtMap)
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

// StringBuilder builds the ASCII art string from input text, colorizing substrings if specified.
func StringBuilder(inputText, subString, colour string, asciiArtMap map[rune]string) string {
	var result strings.Builder

	for i := 0; i < 8; i++ {
		start := 0

		for start < len(inputText) {
			if subString == "" {
				result.WriteString(processNormal(inputText, colour, asciiArtMap, i))
				break
			} else if strings.HasPrefix(inputText[start:], subString) {
				result.WriteString(colorizeSubstring(subString, colour, asciiArtMap, i))
				start += len(subString)
			} else {
				result.WriteString(processCharacter(rune(inputText[start]), asciiArtMap, i))
				start++
			}
		}

		result.WriteString("\n")
	}

	return result.String()
}

// processNormal processes the input text normally, optionally colorizing each character.
func processNormal(inputText, colour string, asciiArtMap map[rune]string, lineIndex int) string {
	var result strings.Builder
	for _, v := range inputText {
		artLines := strings.Split(asciiArtMap[v], "\n")
		if colour != "" {
			ansiCode, _ := color.SetColor(colour)
			result.WriteString(color.Colorize(ansiCode, artLines[lineIndex]))
		} else {
			result.WriteString(artLines[lineIndex])
		}
	}
	return result.String()
}

// colorizeSubstring colorizes the specified substring.
func colorizeSubstring(subString, colour string, asciiArtMap map[rune]string, lineIndex int) string {
	var result strings.Builder
	for _, v := range subString {
		artLines := strings.Split(asciiArtMap[v], "\n")
		ansiCode, _ := color.SetColor(colour)
		result.WriteString(color.Colorize(ansiCode, artLines[lineIndex]))
	}
	return result.String()
}

// processCharacter processes a single character, adding its ASCII art lines.
func processCharacter(char rune, asciiArtMap map[rune]string, lineIndex int) string {
	artLines := strings.Split(asciiArtMap[char], "\n")
	return artLines[lineIndex]
}
