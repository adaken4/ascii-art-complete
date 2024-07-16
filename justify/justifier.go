package justify

import (
	"fmt"
	"strings"
)

func ArtAligner(position, artText string) (string, error) {
	if position == "left" {
		return artText, nil
	}

	var result strings.Builder

	terminalWidth, err := getTerminalWidth()
	if err != nil {
		return "", err
	}
	// Split artText into lines
	lines := strings.Split(artText, "\n")

	spaceIndexes := SpaceIndexes(lines[3], "      ")
	fmt.Println(spaceIndexes)

	// Define the padding
	var paddedLines []string

	// todo
	// calculate num spACES

	// Pad each line individually
	for _, line := range lines {
		textWidth := len(line)
		padding := 0
		switch position {
		case "center":
			padding = (terminalWidth - textWidth) / 2
			paddedLine := strings.Repeat(" ", padding) + line
			paddedLines = append(paddedLines, paddedLine)
		case "right":
			padding = (terminalWidth - textWidth)
			paddedLine := strings.Repeat(" ", padding) + line
			paddedLines = append(paddedLines, paddedLine)
		case "justify":

			padding = (terminalWidth - textWidth) / len(spaceIndexes)
			// fmt.Println(len(line))
			paddedLine := line[:spaceIndexes[0]] + strings.Repeat(" ", padding) + line[spaceIndexes[0]+6:]
			fmt.Println(paddedLine)
			paddedLines = append(paddedLines, paddedLine)
			// TODO
			// each space padding / no.of spaces

		}
	}

	// Write each padded line and ensure the prompt is not padded
	for i, paddedLine := range paddedLines {
		if i == len(paddedLines)-1 {
			result.WriteString("")
		} else {
			result.WriteString(paddedLine + "\n")
		}
	}
	return result.String(), nil
}

func SpaceIndexes(line, spaces string) []int {
	var indexes []int
	for i := 0; i < len(line); {
		index := strings.Index(line[i:], spaces)
		if index == -1 {
			break
		}
		indexes = append(indexes, i+index)
		i += index + 6
	}
	return indexes
}
