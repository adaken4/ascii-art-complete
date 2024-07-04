package justify

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ArtAligner(position, artText string) string {

	var result strings.Builder

	terminalWidth, _ := getTerminalWidth()

	// Split artText into lines
	lines := strings.Split(artText, "\n")

	// Define the padding
	var paddedLines []string

	// Pad each line individually
	for _, line := range lines {
		textWidth := len(line)
		padding := 0
		switch position {
		case "center":
			padding = (terminalWidth - textWidth) / 2
		case "right":
			padding = (terminalWidth - textWidth)
		}
		paddedLine := strings.Repeat(" ", padding) + line
		paddedLines = append(paddedLines, paddedLine)
	}

	// Print each padded line and ensure the prompt is not padded
	for i, paddedLine := range paddedLines {
		if i == len(paddedLines)-1 {
			result.WriteString("")
		} else {
			result.WriteString(paddedLine + "\n")
		}
	}
	return result.String()
}

func getTerminalWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	widthBytes, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	width, err := strconv.Atoi(strings.TrimSpace(string(widthBytes)))
	if err != nil {
		return 0, err
	}

	return width, nil
}
