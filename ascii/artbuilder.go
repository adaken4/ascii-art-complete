package ascii

import (
	"regexp"
	"strings"

	"asciiart/color"
)

type ArtParams struct {
	InputText   string
	SubString   string
	Colour      string
	AsciiArtMap map[rune]string
}

// ArtStringBuilder generates ASCII art for a given string using a specified banner file.
func ArtStringBuilder(params ArtParams) (string, error) {
	var result strings.Builder
	newlines := regexp.MustCompile(`\\n`)
	params.InputText = newlines.ReplaceAllString(params.InputText, "\n")

	// Handle newlines accordingly, if input text contains only newlines
	if onlyNewLines(params.InputText) {
		result.WriteString(params.InputText)
		return result.String(), nil
	}

	inputSlices := strings.Split(params.InputText, "\n")

	for _, v := range inputSlices {
		if v == "" {
			result.WriteString("\n")
		} else {
			artString, err := StringBuilder(ArtParams{
				InputText:   v,
				SubString:   params.SubString,
				Colour:      params.Colour,
				AsciiArtMap: params.AsciiArtMap,
			})
			if err != nil {
				return "", err
			}
			result.WriteString(artString)
		}
	}

	// Return the generated ASCII art string
	return result.String(), nil
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
func StringBuilder(params ArtParams) (string, error) {
	var result strings.Builder

	for i := 0; i < 8; i++ {
		start := 0

		for start < len(params.InputText) {
			if params.SubString == "" {
				normalString, err := processNormal(params, i)
				if err != nil {
					return "", err
				}
				result.WriteString(normalString)
				break
			} else if strings.HasPrefix(params.InputText[start:], params.SubString) {
				coloredSubstring, err := colorizeSubstring(params, i)
				if err != nil {
					return "", err
				}
				result.WriteString(coloredSubstring)
				start += len(params.SubString)
			} else {
				result.WriteString(ProcessCharacter(rune(params.InputText[start]), params.AsciiArtMap, i))
				start++
			}
		}

		result.WriteString("\n")
	}

	return result.String(), nil
}

// processNormal processes the input text normally, optionally colorizing each character
func processNormal(params ArtParams, lineIndex int) (string, error) {
	return processText(params, lineIndex, false)
}

// colorizeSubstring colorizes the specified substring.
func colorizeSubstring(params ArtParams, lineIndex int) (string, error) {
	return processText(params, lineIndex, true)
}

// processCharacter processes a single character, adding its ASCII art lines.
func ProcessCharacter(char rune, asciiArtMap map[rune]string, lineIndex int) string {
	artLines := strings.Split(asciiArtMap[char], "\n")
	return artLines[lineIndex]
}

// processText processes the input text, optionally colorizing it.
func processText(params ArtParams, lineIndex int, isSubstring bool) (string, error) {
	var result strings.Builder
	text := params.InputText
	if isSubstring {
		text = params.SubString
	}

	for _, v := range text {
		artLines := strings.Split(params.AsciiArtMap[v], "\n")
		result.WriteString(artLines[lineIndex])
	}

	if params.Colour != "" {
		ansiCode, err := color.SetColor(params.Colour)
		if err != nil {
			return "", err
		}
		return color.Colorize(ansiCode, result.String()), nil
	}
	return result.String(), nil
}
