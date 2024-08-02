package justify

import (
	"errors"
	"strings"

	"asciiart/ascii"
	"asciiart/color"
)

var differentParams ascii.ArtParams

func ArtAligner(position string, params ascii.ArtParams) (string, error) {
	var result strings.Builder
	differentParams.InputText = params.InputText
	differentParams.AsciiArtMap = params.AsciiArtMap

	uncoloredText, err := ascii.ArtStringBuilder(differentParams)
	if err != nil {
		return "", err
	}

	lines := strings.Split(uncoloredText, "\n")
	lines = lines[:len(lines)-1]
	paddedLines := []string{}

	terminalWidth, err := getTerminalWidth()
	if err != nil {
		return "", err
	}

	switch position {
	case "left":
		ansiCode, _ := color.SetColor(params.Colour)
		return color.Colorize(ansiCode, uncoloredText), nil
	case "right":
		ansiCode, _ := color.SetColor(params.Colour)
		for _, line := range lines {
			artWidth := len(line)
			padding := terminalWidth - artWidth
			paddedLine := strings.Repeat(" ", padding) + color.Colorize(ansiCode, line)
			paddedLines = append(paddedLines, paddedLine)
		}
	case "center":
		ansiCode, _ := color.SetColor(params.Colour)
		for _, line := range lines {
			artWidth := len(line)
			padding := (terminalWidth - artWidth) / 2
			paddedLine := strings.Repeat(" ", padding) + color.Colorize(ansiCode, line)
			paddedLines = append(paddedLines, paddedLine)
		}
	case "justify":
		ansiCode, _ := color.SetColor(params.Colour)
		inputTextLines := strings.Split(params.InputText, "\n")
		for _, line := range inputTextLines {
			// Count spaces in params.InputText
			differentParams.InputText = line
			spaceCount := strings.Count(line, " ")

			lineArtText, _ := ascii.ArtStringBuilder(differentParams)

			linesOfLineArtText := strings.Split(lineArtText, "\n")
			artWidth := len(linesOfLineArtText[0])
			padding := 0
			if spaceCount == 0 {
				padding = 0
			} else if artWidth < terminalWidth {
				padding = (terminalWidth - artWidth) / spaceCount
			} else {
				return "", errors.New("error: terminal size too small, expand terminal")
			}
			paddedLine := ""

			for i := 0; i < 8; i++ {
				for start := 0; start < len(differentParams.InputText); start++ {
					if strings.HasPrefix(differentParams.InputText[start:], " ") {
						paddedLine += strings.Repeat(" ", padding) + "      "

					} else {
						paddedLine += ascii.ProcessCharacter(rune(differentParams.InputText[start]), params.AsciiArtMap, i)

					}
				}
				paddedLine = color.Colorize(ansiCode, paddedLine)
				paddedLines = append(paddedLines, paddedLine)
				paddedLine = ""
			}
		}

	}

	// Join padded lines into result
	for _, v := range paddedLines {
		result.WriteString(v + "\n")
	}

	return result.String(), nil
}
