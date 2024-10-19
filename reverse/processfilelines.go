package reverse

import "strings"

// ProcessReverseFileLines split input string into lines and checks for and removes any custom delimiter "$"
// from the end of each line, and returns the processed lines as a slice of strings
func ProcessReverseFileLines(fileContent string) []string {
	lines := strings.Split(fileContent, "\n")
	processedLines := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSuffix(line, "$")
		processedLines = append(processedLines, line)
	}
	return processedLines
}
